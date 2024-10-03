// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";

import "./IGatewayEVM.sol";
import "./IERC20CustodyNew.sol";

/// @title ERC20CustodyNew
/// @notice Holds the ERC20 tokens deposited on ZetaChain and includes functionality to call a contract.
/// @dev This contract does not call smart contracts directly, it passes through the Gateway contract.
contract ERC20CustodyNew is IERC20CustodyNewEvents, IERC20CustodyNewErrors, ReentrancyGuard {
    using SafeERC20 for IERC20;

    /// @notice Gateway contract.
    IGatewayEVM public gateway;
    /// @notice TSS address.
    address public tssAddress;

    /// @notice Only TSS address allowed modifier.
    modifier onlyTSS() {
        if (msg.sender != tssAddress) {
            revert InvalidSender();
        }
        _;
    }

    constructor(address _gateway, address _tssAddress) {
        if (_gateway == address(0) || _tssAddress == address(0)) {
            revert ZeroAddress();
        }
        gateway = IGatewayEVM(_gateway);
        tssAddress = _tssAddress;
    }

    /// @notice Withdraw directly transfers the tokens to the destination address without contract call.
    /// @dev This function can only be called by the TSS address.
    /// @param token Address of the ERC20 token.
    /// @param to Destination address for the tokens.
    /// @param amount Amount of tokens to withdraw.
    function withdraw(address token, address to, uint256 amount) external nonReentrant onlyTSS {
        IERC20(token).safeTransfer(to, amount);

        emit Withdraw(token, to, amount);
    }

    /// @notice WithdrawAndCall transfers tokens to Gateway and call a contract through the Gateway.
    /// @dev This function can only be called by the TSS address.
    /// @param token Address of the ERC20 token.
    /// @param to Address of the contract to call.
    /// @param amount Amount of tokens to withdraw.
    /// @param data Calldata to pass to the contract call.
    function withdrawAndCall(address token, address to, uint256 amount, bytes calldata data) public nonReentrant onlyTSS {
        // Transfer the tokens to the Gateway contract
        IERC20(token).safeTransfer(address(gateway), amount);

        // Forward the call to the Gateway contract
        gateway.executeWithERC20(token, to, amount, data);

        emit WithdrawAndCall(token, to, amount, data);
    }

    /// @notice WithdrawAndRevert transfers tokens to Gateway and call a contract with a revert functionality through the Gateway.
    /// @dev This function can only be called by the TSS address.
    /// @param token Address of the ERC20 token.
    /// @param to Address of the contract to call.
    /// @param amount Amount of tokens to withdraw.
    /// @param data Calldata to pass to the contract call.
    function withdrawAndRevert(address token, address to, uint256 amount, bytes calldata data) public nonReentrant onlyTSS {
        // Transfer the tokens to the Gateway contract
        IERC20(token).safeTransfer(address(gateway), amount);

        // Forward the call to the Gateway contract
        gateway.revertWithERC20(token, to, amount, data);

        emit WithdrawAndRevert(token, to, amount, data);
    }
}