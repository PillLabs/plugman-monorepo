// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";

import "./IGatewayEVM.sol";
import "./IZetaConnector.sol";

/// @title ZetaConnectorNewBase
/// @notice Abstract base contract for ZetaConnector.
/// @dev This contract implements basic functionality for handling tokens and interacting with the Gateway contract.
abstract contract ZetaConnectorNewBase is IZetaConnectorEvents, ReentrancyGuard {
    using SafeERC20 for IERC20;

    /// @notice Error indicating that a zero address was provided.
    error ZeroAddress();
    /// @notice Error indicating that the sender is invalid.
    error InvalidSender();

    /// @notice The Gateway contract used for executing cross-chain calls.
    IGatewayEVM public immutable gateway;
    /// @notice The address of the Zeta token.
    address public immutable zetaToken;
    /// @notice The address of the TSS (Threshold Signature Scheme) contract.
    address public tssAddress;

    /// @dev Only TSS address allowed modifier.
    modifier onlyTSS() {
        if (msg.sender != tssAddress) {
            revert InvalidSender();
        }
        _;
    }

    constructor(address _gateway, address _zetaToken, address _tssAddress) {
        if (_gateway == address(0) || _zetaToken == address(0) || _tssAddress == address(0)) {
            revert ZeroAddress();
        }
        gateway = IGatewayEVM(_gateway);
        zetaToken = _zetaToken;
        tssAddress = _tssAddress;
    }

    /// @notice Withdraw tokens to a specified address.
    /// @param to The address to withdraw tokens to.
    /// @param amount The amount of tokens to withdraw.
    /// @param internalSendHash A hash used for internal tracking of the transaction.
    function withdraw(address to, uint256 amount, bytes32 internalSendHash) external virtual;

    /// @notice Withdraw tokens and call a contract through Gateway.
    /// @param to The address to withdraw tokens to.
    /// @param amount The amount of tokens to withdraw.
    /// @param data The calldata to pass to the contract call.
    /// @param internalSendHash A hash used for internal tracking of the transaction.
    function withdrawAndCall(address to, uint256 amount, bytes calldata data, bytes32 internalSendHash) external virtual;

    /// @notice Withdraw tokens and call a contract with a revert callback through Gateway.
    /// @param to The address to withdraw tokens to.
    /// @param amount The amount of tokens to withdraw.
    /// @param data The calldata to pass to the contract call.
    /// @param internalSendHash A hash used for internal tracking of the transaction.
    function withdrawAndRevert(address to, uint256 amount, bytes calldata data, bytes32 internalSendHash) external virtual;

    /// @notice Handle received tokens.
    /// @param amount The amount of tokens received.
    function receiveTokens(uint256 amount) external virtual;
}
