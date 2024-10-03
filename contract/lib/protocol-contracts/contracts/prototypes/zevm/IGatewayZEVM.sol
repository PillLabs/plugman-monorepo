// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

import "../../zevm/interfaces/zContract.sol";

/// @title IGatewayZEVM
/// @notice Interface for the GatewayZEVM contract.
/// @dev Defines functions for cross-chain interactions and token handling.
interface IGatewayZEVM {
    /// @notice Withdraw ZRC20 tokens to an external chain.
    /// @param receiver The receiver address on the external chain.
    /// @param amount The amount of tokens to withdraw.
    /// @param zrc20 The address of the ZRC20 token.
    function withdraw(bytes memory receiver, uint256 amount, address zrc20) external;

    /// @notice Withdraw ZRC20 tokens and call a smart contract on an external chain.
    /// @param receiver The receiver address on the external chain.
    /// @param amount The amount of tokens to withdraw.
    /// @param zrc20 The address of the ZRC20 token.
    /// @param message The calldata to pass to the contract call.
    function withdrawAndCall(bytes memory receiver, uint256 amount, address zrc20, bytes calldata message) external;

    /// @notice Call a smart contract on an external chain without asset transfer.
    /// @param receiver The receiver address on the external chain.
    /// @param message The calldata to pass to the contract call.
    function call(bytes memory receiver, bytes calldata message) external;

    /// @notice Deposit foreign coins into ZRC20.
    /// @param zrc20 The address of the ZRC20 token.
    /// @param amount The amount of tokens to deposit.
    /// @param target The target address to receive the deposited tokens.
    function deposit(
        address zrc20,
        uint256 amount,
        address target
    ) external;

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
    ) external;

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
    ) external;
}

/// @title IGatewayZEVMEvents
/// @notice Interface for the events emitted by the GatewayZEVM contract.
interface IGatewayZEVMEvents {
    /// @notice Emitted when a cross-chain call is made.
    /// @param sender The address of the sender.
    /// @param receiver The receiver address on the external chain.
    /// @param message The calldata passed to the contract call.
    event Call(address indexed sender, bytes receiver, bytes message);

    /// @notice Emitted when a withdrawal is made.
    /// @param from The address from which the tokens are withdrawn.
    /// @param zrc20 The address of the ZRC20 token.
    /// @param to The receiver address on the external chain.
    /// @param value The amount of tokens withdrawn.
    /// @param gasfee The gas fee for the withdrawal.
    /// @param protocolFlatFee The protocol flat fee for the withdrawal.
    /// @param message The calldata passed to the contract call.
    event Withdrawal(address indexed from, address zrc20, bytes to, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message);
}

/// @title IGatewayZEVMErrors
/// @notice Interface for the errors used in the GatewayZEVM contract.
interface IGatewayZEVMErrors {
    /// @notice Error indicating a withdrawal failure.
    error WithdrawalFailed();

    /// @notice Error indicating an insufficient ZRC20 token amount.
    error InsufficientZRC20Amount();

    /// @notice Error indicating a failure to burn ZRC20 tokens.
    error ZRC20BurnFailed();

    /// @notice Error indicating a failure to transfer ZRC20 tokens.
    error ZRC20TransferFailed();

    /// @notice Error indicating a failure to transfer gas fee.
    error GasFeeTransferFailed();

    /// @notice Error indicating that the caller is not the Fungible module.
    error CallerIsNotFungibleModule();

    /// @notice Error indicating an invalid target address.
    error InvalidTarget();

    /// @notice Error indicating a failure to send ZETA tokens.
    error FailedZetaSent();

    /// @notice Error indicating that only WZETA or the Fungible module can call the function.
    error OnlyWZETAOrFungible();
}