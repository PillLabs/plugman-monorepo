// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/// @title IGatewayEVMEvents
/// @notice Interface for the events emitted by the GatewayEVM contract.
interface IGatewayEVMEvents {
    /// @notice Emitted when a contract call is executed.
    /// @param destination The address of the contract called.
    /// @param value The amount of ETH sent with the call.
    /// @param data The calldata passed to the contract call.
    event Executed(address indexed destination, uint256 value, bytes data);

    /// @notice Emitted when a contract call is reverted.
    /// @param destination The address of the contract called.
    /// @param value The amount of ETH sent with the call.
    /// @param data The calldata passed to the contract call.
    event Reverted(address indexed destination, uint256 value, bytes data);

    /// @notice Emitted when a contract call with ERC20 tokens is executed.
    /// @param token The address of the ERC20 token.
    /// @param to The address of the contract called.
    /// @param amount The amount of tokens transferred.
    /// @param data The calldata passed to the contract call.
    event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data);

    /// @notice Emitted when a contract call with ERC20 tokens is reverted.
    /// @param token The address of the ERC20 token.
    /// @param to The address of the contract called.
    /// @param amount The amount of tokens transferred.
    /// @param data The calldata passed to the contract call.
    event RevertedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data);

    /// @notice Emitted when a deposit is made.
    /// @param sender The address of the sender.
    /// @param receiver The address of the receiver.
    /// @param amount The amount of ETH or tokens deposited.
    /// @param asset The address of the ERC20 token (zero address if ETH).
    /// @param payload The calldata passed with the deposit.
    event Deposit(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload);

    /// @notice Emitted when an omnichain smart contract call is made without asset transfer.
    /// @param sender The address of the sender.
    /// @param receiver The address of the receiver.
    /// @param payload The calldata passed to the call.
    event Call(address indexed sender, address indexed receiver, bytes payload);
}

/// @title IGatewayEVMErrors
/// @notice Interface for the errors used in the GatewayEVM contract.
interface IGatewayEVMErrors {
    /// @notice Error for failed execution.
    error ExecutionFailed();

    /// @notice Error for failed deposit.
    error DepositFailed();

    /// @notice Error for insufficient ETH amount.
    error InsufficientETHAmount();

    /// @notice Error for insufficient ERC20 token amount.
    error InsufficientERC20Amount();

    /// @notice Error for zero address input.
    error ZeroAddress();

    /// @notice Error for failed token approval.
    error ApprovalFailed();

    /// @notice Error for already initialized custody.
    error CustodyInitialized();

    /// @notice Error for invalid sender.
    error InvalidSender();
}

/// @title IGatewayEVM
/// @notice Interface for the GatewayEVM contract.
interface IGatewayEVM {
    /// @notice Executes a call to a contract using ERC20 tokens.
    /// @param token The address of the ERC20 token.
    /// @param to The address of the contract to call.
    /// @param amount The amount of tokens to transfer.
    /// @param data The calldata to pass to the contract call.
    function executeWithERC20(
        address token,
        address to,
        uint256 amount,
        bytes calldata data
    ) external;

    /// @notice Executes a call to a contract.
    /// @param destination The address of the contract to call.
    /// @param data The calldata to pass to the contract call.
    /// @return The result of the contract call.
    function execute(address destination, bytes calldata data) external payable returns (bytes memory);

    /// @notice Executes a revertable call to a contract using ERC20 tokens.
    /// @param token The address of the ERC20 token.
    /// @param to The address of the contract to call.
    /// @param amount The amount of tokens to transfer.
    /// @param data The calldata to pass to the contract call.
    function revertWithERC20(
        address token,
        address to,
        uint256 amount,
        bytes calldata data
    ) external;
}

/// @title Revertable
/// @notice Interface for contracts that support revertable calls.
interface Revertable {
    /// @notice Called when a revertable call is made.
    /// @param data The calldata to pass to the revertable call.
    function onRevert(bytes calldata data) external;
}
