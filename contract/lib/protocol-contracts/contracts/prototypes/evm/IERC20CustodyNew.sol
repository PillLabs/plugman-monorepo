// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/// @title IERC20CustodyNewEvents
/// @notice Interface for the events emitted by the ERC20 custody contract.
interface IERC20CustodyNewEvents {
    /// @notice Emitted when tokens are withdrawn.
    /// @param token The address of the ERC20 token.
    /// @param to The address receiving the tokens.
    /// @param amount The amount of tokens withdrawn.
    event Withdraw(address indexed token, address indexed to, uint256 amount);

    /// @notice Emitted when tokens are withdrawn and a contract call is made.
    /// @param token The address of the ERC20 token.
    /// @param to The address receiving the tokens.
    /// @param amount The amount of tokens withdrawn.
    /// @param data The calldata passed to the contract call.
    event WithdrawAndCall(address indexed token, address indexed to, uint256 amount, bytes data);

    /// @notice Emitted when tokens are withdrawn and a revertable contract call is made.
    /// @param token The address of the ERC20 token.
    /// @param to The address receiving the tokens.
    /// @param amount The amount of tokens withdrawn.
    /// @param data The calldata passed to the contract call.
    event WithdrawAndRevert(address indexed token, address indexed to, uint256 amount, bytes data);
}

/// @title IERC20CustodyNewErrors
/// @notice Interface for the errors used in the ERC20 custody contract.
interface IERC20CustodyNewErrors {
    /// @notice Error for zero address input.
    error ZeroAddress();

    /// @notice Error for invalid sender.
    error InvalidSender();
}