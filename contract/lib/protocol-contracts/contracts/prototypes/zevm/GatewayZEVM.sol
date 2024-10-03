// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "../../zevm/interfaces/IZRC20.sol";
import "../../zevm/interfaces/zContract.sol";
import "./IGatewayZEVM.sol";
import "../../zevm/interfaces/IWZETA.sol";

/// @title GatewayZEVM
/// @notice The GatewayZEVM contract is the endpoint to call smart contracts on omnichain.
/// @dev The contract doesn't hold any funds and should never have active allowances.
contract GatewayZEVM is IGatewayZEVMEvents, IGatewayZEVMErrors, Initializable, OwnableUpgradeable, UUPSUpgradeable, ReentrancyGuardUpgradeable {
    /// @notice Error indicating a zero address was provided.
    error ZeroAddress();

    /// @notice The constant address of the Fungible module.
    address public constant FUNGIBLE_MODULE_ADDRESS = 0x735b14BB79463307AAcBED86DAf3322B1e6226aB;
    /// @notice The address of the Zeta token.
    address public zetaToken;

    /// @dev Only Fungible module address allowed modifier.
    modifier onlyFungible() {
        if (msg.sender != FUNGIBLE_MODULE_ADDRESS) {
            revert CallerIsNotFungibleModule();
        }
        _;
    }

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize(address _zetaToken) public initializer {
        if (_zetaToken == address(0)) {
            revert ZeroAddress();
        }
    
        __Ownable_init();
        __UUPSUpgradeable_init();
        __ReentrancyGuard_init();
        zetaToken = _zetaToken;
    }

    /// @dev Authorizes the upgrade of the contract.
    /// @param newImplementation The address of the new implementation.
    function _authorizeUpgrade(address newImplementation) internal override onlyOwner() {}

    /// @dev Receive function to receive ZETA from WETH9.withdraw().
    receive() external payable {
        if (msg.sender != zetaToken && msg.sender != FUNGIBLE_MODULE_ADDRESS) revert OnlyWZETAOrFungible();
    }

    /// @dev Internal function to withdraw ZRC20 tokens.
    /// @param amount The amount of tokens to withdraw.
    /// @param zrc20 The address of the ZRC20 token.
    /// @return The gas fee for the withdrawal.
    function _withdrawZRC20(uint256 amount, address zrc20) internal returns (uint256) {
        (address gasZRC20, uint256 gasFee) = IZRC20(zrc20).withdrawGasFee();
        if (!IZRC20(gasZRC20).transferFrom(msg.sender, FUNGIBLE_MODULE_ADDRESS, gasFee)) {
            revert GasFeeTransferFailed();
        }

        if (!IZRC20(zrc20).transferFrom(msg.sender, address(this), amount)) {
            revert ZRC20TransferFailed();
        }

        if (!IZRC20(zrc20).burn(amount)) revert ZRC20BurnFailed();

        return gasFee;
    }

    /// @dev Internal function to transfer ZETA tokens.
    /// @param amount The amount of tokens to transfer.
    /// @param to The address to transfer the tokens to.
    function _transferZETA(uint256 amount, address to) internal {
        if (!IWETH9(zetaToken).transferFrom(msg.sender, address(this), amount)) revert FailedZetaSent();
        IWETH9(zetaToken).withdraw(amount);
        (bool sent, ) = to.call{value: amount}("");
        if (!sent) revert FailedZetaSent();
    }

    /// @notice Withdraw ZRC20 tokens to an external chain.
    /// @param receiver The receiver address on the external chain.
    /// @param amount The amount of tokens to withdraw.
    /// @param zrc20 The address of the ZRC20 token.
    function withdraw(bytes memory receiver, uint256 amount, address zrc20) external nonReentrant {
        uint256 gasFee = _withdrawZRC20(amount, zrc20);
        emit Withdrawal(msg.sender, zrc20, receiver, amount, gasFee, IZRC20(zrc20).PROTOCOL_FLAT_FEE(), "");
    }

    /// @notice Withdraw ZRC20 tokens and call a smart contract on an external chain.
    /// @param receiver The receiver address on the external chain.
    /// @param amount The amount of tokens to withdraw.
    /// @param zrc20 The address of the ZRC20 token.
    /// @param message The calldata to pass to the contract call.
    function withdrawAndCall(bytes memory receiver, uint256 amount, address zrc20, bytes calldata message) external nonReentrant {
        uint256 gasFee = _withdrawZRC20(amount, zrc20);
        emit Withdrawal(msg.sender, zrc20, receiver, amount, gasFee, IZRC20(zrc20).PROTOCOL_FLAT_FEE(), message);
    }

    /// @notice Withdraw ZETA tokens to an external chain.
    /// @param amount The amount of tokens to withdraw.
    function withdraw(uint256 amount) external nonReentrant {
        _transferZETA(amount, FUNGIBLE_MODULE_ADDRESS);
        emit Withdrawal(msg.sender, address(zetaToken), abi.encodePacked(FUNGIBLE_MODULE_ADDRESS), amount, 0, 0, "");
    }

    /// @notice Withdraw ZETA tokens and call a smart contract on an external chain.
    /// @param amount The amount of tokens to withdraw.
    /// @param message The calldata to pass to the contract call.
    function withdrawAndCall(uint256 amount, bytes calldata message) external nonReentrant {
        _transferZETA(amount, FUNGIBLE_MODULE_ADDRESS);
        emit Withdrawal(msg.sender, address(zetaToken), abi.encodePacked(FUNGIBLE_MODULE_ADDRESS), amount, 0, 0, message);
    }

    /// @notice Call a smart contract on an external chain without asset transfer.
    /// @param receiver The receiver address on the external chain.
    /// @param message The calldata to pass to the contract call.
    function call(bytes memory receiver, bytes calldata message) external nonReentrant {
        emit Call(msg.sender, receiver, message);
    }

    /// @notice Deposit foreign coins into ZRC20.
    /// @param zrc20 The address of the ZRC20 token.
    /// @param amount The amount of tokens to deposit.
    /// @param target The target address to receive the deposited tokens.
    function deposit(
        address zrc20,
        uint256 amount,
        address target
    ) external onlyFungible {
        if (target == FUNGIBLE_MODULE_ADDRESS || target == address(this)) revert InvalidTarget();

        IZRC20(zrc20).deposit(target, amount);
    }

    /// @notice Execute a user-specified contract on ZEVM.
    /// @param context The context of the cross-chain call.
    /// @param zrc20 The address of the ZRC20 token.
    /// @param amount The amount of tokens to transfer.
    /// @param target The target contract to call.
    /// @param message The calldata to pass to the contract call.
    function execute(
        zContext calldata context,
        address zrc20,
        uint256 amount,
        address target,
        bytes calldata message
    ) external onlyFungible {
        UniversalContract(target).onCrossChainCall(context, zrc20, amount, message);
    }

    /// @notice Deposit foreign coins into ZRC20 and call a user-specified contract on ZEVM.
    /// @param context The context of the cross-chain call.
    /// @param zrc20 The address of the ZRC20 token.
    /// @param amount The amount of tokens to transfer.
    /// @param target The target contract to call.
    /// @param message The calldata to pass to the contract call.
    function depositAndCall(
        zContext calldata context,
        address zrc20,
        uint256 amount,
        address target,
        bytes calldata message
    ) external onlyFungible {
        if (target == FUNGIBLE_MODULE_ADDRESS || target == address(this)) revert InvalidTarget();

        IZRC20(zrc20).deposit(target, amount);
        UniversalContract(target).onCrossChainCall(context, zrc20, amount, message);
    }

    /// @notice Deposit ZETA and call a user-specified contract on ZEVM.
    /// @param context The context of the cross-chain call.
    /// @param amount The amount of tokens to transfer.
    /// @param target The target contract to call.
    /// @param message The calldata to pass to the contract call.
    function depositAndCall(
        zContext calldata context,
        uint256 amount,
        address target,
        bytes calldata message
    ) external onlyFungible {
        if (target == FUNGIBLE_MODULE_ADDRESS || target == address(this)) revert InvalidTarget();

        _transferZETA(amount, target);
        UniversalContract(target).onCrossChainCall(context, zetaToken, amount, message);
    }

    /// @notice Revert a user-specified contract on ZEVM.
    /// @param context The context of the revert call.
    /// @param zrc20 The address of the ZRC20 token.
    /// @param amount The amount of tokens to revert.
    /// @param target The target contract to call.
    /// @param message The calldata to pass to the contract call.
    function executeRevert(
        revertContext calldata context,
        address zrc20,
        uint256 amount,
        address target,
        bytes calldata message
    ) external onlyFungible {
        UniversalContract(target).onRevert(context, zrc20, amount, message);
    }

    /// @notice Deposit foreign coins into ZRC20 and revert a user-specified contract on ZEVM.
    /// @param context The context of the revert call.
    /// @param zrc20 The address of the ZRC20 token.
    /// @param amount The amount of tokens to revert.
    /// @param target The target contract to call.
    /// @param message The calldata to pass to the contract call.
    function depositAndRevert(
        revertContext calldata context,
        address zrc20,
        uint256 amount,
        address target,
        bytes calldata message
    ) external onlyFungible {
        if (target == FUNGIBLE_MODULE_ADDRESS || target == address(this)) revert InvalidTarget();

        IZRC20(zrc20).deposit(target, amount);
        UniversalContract(target).onRevert(context, zrc20, amount, message);
    }
}
