// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./IGatewayZEVM.sol";
import "../../zevm/interfaces/IZRC20.sol";

/// @title SenderZEVM
/// @notice This contract is used just for testing.
/// @dev Provides functions to call a receiver on EVM and to withdraw and call a receiver on EVM.
contract SenderZEVM {
    /// @notice The address of the gateway contract.
    address public gateway;

    /// @notice Error indicating that the approval of tokens failed.
    error ApprovalFailed();

    constructor(address _gateway) {
        gateway = _gateway;
    }

    /// @notice Call a receiver on EVM.
    /// @param receiver The address of the receiver on the external chain.
    /// @param str A string parameter to pass to the receiver's function.
    /// @param num A numeric parameter to pass to the receiver's function.
    /// @param flag A boolean parameter to pass to the receiver's function.
    /// @dev Encodes the function call and passes it to the gateway.
    function callReceiver(bytes memory receiver, string memory str, uint256 num, bool flag) external {
        // Encode the function call to the receiver's receivePayable method
        bytes memory message = abi.encodeWithSignature("receivePayable(string,uint256,bool)", str, num, flag);

        // Pass encoded call to gateway
        IGatewayZEVM(gateway).call(receiver, message);
    }

    /// @notice Withdraw and call a receiver on EVM.
    /// @param receiver The address of the receiver on the external chain.
    /// @param amount The amount of tokens to withdraw.
    /// @param zrc20 The address of the ZRC20 token.
    /// @param str A string parameter to pass to the receiver's function.
    /// @param num A numeric parameter to pass to the receiver's function.
    /// @param flag A boolean parameter to pass to the receiver's function.
    /// @dev Approves the gateway to withdraw tokens and encodes the function call to pass to the gateway.
    function withdrawAndCallReceiver(bytes memory receiver, uint256 amount, address zrc20, string memory str, uint256 num, bool flag) external {
        // Encode the function call to the receiver's receivePayable method
        bytes memory message = abi.encodeWithSignature("receivePayable(string,uint256,bool)", str, num, flag);

        // Approve gateway to withdraw
        if (!IZRC20(zrc20).approve(gateway, amount)) revert ApprovalFailed();

        // Pass encoded call to gateway
        IGatewayZEVM(gateway).withdrawAndCall(receiver, amount, zrc20, message);
    }
}
