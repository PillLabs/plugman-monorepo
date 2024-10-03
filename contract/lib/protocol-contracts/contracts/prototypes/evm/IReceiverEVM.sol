// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/// @title IReceiverEVMEvents
/// @notice Interface for the events emitted by the ReceiverEVM contract.
interface IReceiverEVMEvents {
    /// @notice Emitted when a payable function is called.
    /// @param sender The address of the sender.
    /// @param value The amount of ETH sent with the call.
    /// @param str A string parameter.
    /// @param num A numeric parameter.
    /// @param flag A boolean parameter.
    event ReceivedPayable(address sender, uint256 value, string str, uint256 num, bool flag);

    /// @notice Emitted when a non-payable function is called.
    /// @param sender The address of the sender.
    /// @param strs An array of string parameters.
    /// @param nums An array of numeric parameters.
    /// @param flag A boolean parameter.
    event ReceivedNonPayable(address sender, string[] strs, uint256[] nums, bool flag);

    /// @notice Emitted when ERC20 tokens are received.
    /// @param sender The address of the sender.
    /// @param amount The amount of ERC20 tokens received.
    /// @param token The address of the ERC20 token.
    /// @param destination The address to which the tokens are sent.
    event ReceivedERC20(address sender, uint256 amount, address token, address destination);

    /// @notice Emitted when a function without parameters is called.
    /// @param sender The address of the sender.
    event ReceivedNoParams(address sender);

    /// @notice Emitted when a revert callback function is called.
    /// @param sender The address of the sender.
    /// @param data The calldata passed during the revert.
    event ReceivedRevert(address sender, bytes data);
}
