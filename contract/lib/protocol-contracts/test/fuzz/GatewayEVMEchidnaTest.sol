// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../../contracts/prototypes/evm/GatewayEVM.sol";
import "../../contracts/prototypes/evm/TestERC20.sol";
import "../../contracts/prototypes/evm/ERC20CustodyNew.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

contract GatewayEVMEchidnaTest is GatewayEVM {
    using SafeERC20 for IERC20;

    TestERC20 public testERC20;
    address public echidnaCaller = msg.sender;

    constructor() {
        tssAddress = echidnaCaller;
        zetaConnector = address(0x123);
        testERC20 = new TestERC20("test", "TEST");
        custody = address(new ERC20CustodyNew(address(this), tssAddress));
    }

    function testExecuteWithERC20(address to, uint256 amount, bytes calldata data) public {
        testERC20.mint(address(this), amount);

        executeWithERC20(address(testERC20), to, amount, data);

        // Assertion to ensure no ERC20 tokens are left in the contract after execution
        assert(testERC20.balanceOf(address(this)) == 0);
    }
}