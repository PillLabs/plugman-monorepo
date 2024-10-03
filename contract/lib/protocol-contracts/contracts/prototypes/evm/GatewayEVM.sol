// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";

import "./IGatewayEVM.sol";
import "./ZetaConnectorNewBase.sol";

/// @title GatewayEVM
/// @notice The GatewayEVM contract is the endpoint to call smart contracts on external chains.
/// @dev The contract doesn't hold any funds and should never have active allowances.
contract GatewayEVM is Initializable, OwnableUpgradeable, UUPSUpgradeable, IGatewayEVMErrors, IGatewayEVMEvents, ReentrancyGuardUpgradeable {
    using SafeERC20 for IERC20;

    /// @notice The address of the custody contract.
    address public custody;
    /// @notice The address of the TSS (Threshold Signature Scheme) contract.
    address public tssAddress;
    /// @notice The address of the ZetaConnector contract.
    address public zetaConnector;
    /// @notice The address of the Zeta token contract.
    address public zetaToken;

    /// @notice Only TSS address allowed modifier.
    modifier onlyTSS() {
        if (msg.sender != tssAddress) {
            revert InvalidSender();
        }
        _;
    }

    /// @notice Only custody or connector address allowed modifier.
    modifier onlyAssetHandler() {
        if (msg.sender != custody && msg.sender != zetaConnector) {
            revert InvalidSender();
        }
        _;
    }

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize(address _tssAddress, address _zetaToken) public initializer {
        if (_tssAddress == address(0) || _zetaToken == address(0)) {
            revert ZeroAddress();
        }
        
        __Ownable_init();
        __UUPSUpgradeable_init();
        __ReentrancyGuard_init();

        tssAddress = _tssAddress;
        zetaToken = _zetaToken;
    }

    /// @dev Authorizes the upgrade of the contract, sender must be owner.
    /// @param newImplementation Address of the new implementation.
    function _authorizeUpgrade(address newImplementation) internal override onlyOwner() {}

    /// @dev Internal function to execute a call to a destination address.
    /// @param destination Address to call.
    /// @param data Calldata to pass to the call.
    /// @return The result of the call.
    function _execute(address destination, bytes calldata data) internal returns (bytes memory) {
        (bool success, bytes memory result) = destination.call{value: msg.value}(data);
        if (!success) revert ExecutionFailed();

        return result;
    }

    /// @notice Transfers msg.value to destination contract and executes it's onRevert function.
    /// @dev This function can only be called by the TSS address and it is payable.
    /// @param destination Address to call.
    /// @param data Calldata to pass to the call.
    function executeRevert(address destination, bytes calldata data) public payable onlyTSS {
        (bool success, bytes memory result) = destination.call{value: msg.value}("");
        if (!success) revert ExecutionFailed();
        Revertable(destination).onRevert(data);
        
        emit Reverted(destination, msg.value, data);
    }

    /// @notice Executes a call to a destination address without ERC20 tokens.
    /// @dev This function can only be called by the TSS address and it is payable.
    /// @param destination Address to call.
    /// @param data Calldata to pass to the call.
    /// @return The result of the call.
    function execute(address destination, bytes calldata data) external payable onlyTSS returns (bytes memory)  {
        bytes memory result = _execute(destination, data);

        emit Executed(destination, msg.value, data);

        return result;
    }

    /// @notice Executes a call to a destination contract using ERC20 tokens.
    /// @dev This function can only be called by the custody or connector address.
    ///      It uses the ERC20 allowance system, resetting gateway allowance at the end.
    /// @param token Address of the ERC20 token.
    /// @param to Address of the contract to call.
    /// @param amount Amount of tokens to transfer.
    /// @param data Calldata to pass to the call.
    function executeWithERC20(
        address token,
        address to,
        uint256 amount,
        bytes calldata data
    ) public nonReentrant onlyAssetHandler {
        if (amount == 0) revert InsufficientERC20Amount();
        // Approve the target contract to spend the tokens
        if(!resetApproval(token, to)) revert ApprovalFailed();
        if(!IERC20(token).approve(to, amount)) revert ApprovalFailed();
        // Execute the call on the target contract
        bytes memory result = _execute(to, data);

        // Reset approval
        if(!resetApproval(token, to)) revert ApprovalFailed();

        // Transfer any remaining tokens back to the custody/connector contract
        uint256 remainingBalance = IERC20(token).balanceOf(address(this));
        if (remainingBalance > 0) {
            transferToAssetHandler(token, remainingBalance);
        }

        emit ExecutedWithERC20(token, to, amount, data);
    }

    /// @notice Directly transfers ERC20 tokens and calls onRevert.
    /// @dev This function can only be called by the custody or connector address.
    /// @param token Address of the ERC20 token.
    /// @param to Address of the contract to call.
    /// @param amount Amount of tokens to transfer.
    /// @param data Calldata to pass to the call.
    function revertWithERC20(
        address token,
        address to,
        uint256 amount,
        bytes calldata data
    ) external nonReentrant onlyAssetHandler {
        if (amount == 0) revert InsufficientERC20Amount();

        IERC20(token).safeTransfer(address(to), amount);
        Revertable(to).onRevert(data);

        emit RevertedWithERC20(token, to, amount, data);
    }

    /// @notice Deposits ETH to the TSS address.
    /// @param receiver Address of the receiver.
    function deposit(address receiver) external payable {
        if (msg.value == 0) revert InsufficientETHAmount();
        (bool deposited, ) = tssAddress.call{value: msg.value}("");

        if (deposited == false) revert DepositFailed();
        
        emit Deposit(msg.sender, receiver, msg.value, address(0), "");
    }

    /// @notice Deposits ERC20 tokens to the custody or connector contract.
    /// @param receiver Address of the receiver.
    /// @param amount Amount of tokens to deposit.
    /// @param asset Address of the ERC20 token.
    function deposit(address receiver, uint256 amount, address asset) external {
        if (amount == 0) revert InsufficientERC20Amount();

        transferFromToAssetHandler(msg.sender, asset, amount);

        emit Deposit(msg.sender, receiver, amount, asset, "");
    }

    /// @notice Deposits ETH to the TSS address and calls an omnichain smart contract.
    /// @param receiver Address of the receiver.
    /// @param payload Calldata to pass to the call.
    function depositAndCall(address receiver, bytes calldata payload) external payable {
        if (msg.value == 0) revert InsufficientETHAmount();
        (bool deposited, ) = tssAddress.call{value: msg.value}("");

        if (deposited == false) revert DepositFailed();
        
        emit Deposit(msg.sender, receiver, msg.value, address(0), payload);
    }

    /// @notice Deposits ERC20 tokens to the custody or connector contract and calls an omnichain smart contract.
    /// @param receiver Address of the receiver.
    /// @param amount Amount of tokens to deposit.
    /// @param asset Address of the ERC20 token.
    /// @param payload Calldata to pass to the call.
    function depositAndCall(address receiver, uint256 amount, address asset, bytes calldata payload) external {
        if (amount == 0) revert InsufficientERC20Amount();
       
        transferFromToAssetHandler(msg.sender, asset, amount);

        emit Deposit(msg.sender, receiver, amount, asset, payload);
    }

    /// @notice Calls an omnichain smart contract without asset transfer.
    /// @param receiver Address of the receiver.
    /// @param payload Calldata to pass to the call.
    function call(address receiver, bytes calldata payload) external {
        emit Call(msg.sender, receiver, payload);
    }

    /// @notice Sets the custody contract address.
    /// @param _custody Address of the custody contract.
    function setCustody(address _custody) external onlyTSS {
        if (custody != address(0)) revert CustodyInitialized();
        if (_custody == address(0)) revert ZeroAddress();

        custody = _custody;
    }

    /// @notice Sets the connector contract address.
    /// @param _zetaConnector Address of the connector contract.
    function setConnector(address _zetaConnector) external onlyTSS {
        if (zetaConnector != address(0)) revert CustodyInitialized();
        if (_zetaConnector == address(0)) revert ZeroAddress();

        zetaConnector = _zetaConnector;
    }

    /// @dev Resets the approval of a token for a specified address.
    /// This is used to ensure that the approval is set to zero before setting it to a new value.
    /// @param token Address of the ERC20 token.
    /// @param to Address to reset the approval for.
    /// @return True if the approval reset was successful, false otherwise.
    function resetApproval(address token, address to) private returns (bool) {
        return IERC20(token).approve(to, 0);
    }

    /// @dev Transfers tokens from the sender to the asset handler.
    /// This function handles the transfer of tokens to either the connector or custody contract based on the asset type.
    /// @param from Address of the sender.
    /// @param token Address of the ERC20 token.
    /// @param amount Amount of tokens to transfer.
    function transferFromToAssetHandler(address from, address token, uint256 amount) private {
        if (token == zetaToken) { // transfer to connector
            // transfer amount to gateway
            IERC20(token).safeTransferFrom(from, address(this), amount);
            // approve connector to handle tokens depending on connector version (eg. lock or burn)
            IERC20(token).approve(zetaConnector, amount);
            // send tokens to connector
            ZetaConnectorNewBase(zetaConnector).receiveTokens(amount);
        } else { // transfer to custody
            IERC20(token).safeTransferFrom(from, custody, amount);
        }
    }

    /// @dev Transfers tokens to the asset handler.
    /// This function handles the transfer of tokens to either the connector or custody contract based on the asset type.
    /// @param token Address of the ERC20 token.
    /// @param amount Amount of tokens to transfer.
    function transferToAssetHandler(address token, uint256 amount) private {
        if (token == zetaToken) { // transfer to connector
            // approve connector to handle tokens depending on connector version (eg. lock or burn)
            IERC20(token).approve(zetaConnector, amount);
            // send tokens to connector
            ZetaConnectorNewBase(zetaConnector).receiveTokens(amount);
        } else { // transfer to custody
            IERC20(token).safeTransfer(custody, amount);
        }
    }
}
