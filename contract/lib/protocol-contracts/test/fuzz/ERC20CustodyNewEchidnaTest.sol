// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../../contracts/prototypes/evm/TestERC20.sol";
import "../../contracts/prototypes/evm/ERC20CustodyNew.sol";
import "../../contracts/prototypes/evm/GatewayEVM.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {Upgrades} from "openzeppelin-foundry-upgrades/LegacyUpgrades.sol";

contract ERC20CustodyNewEchidnaTest is ERC20CustodyNew {
    using SafeERC20 for IERC20;

    TestERC20 public testERC20;
    address public echidnaCaller = msg.sender;

    address proxy = address(new ERC1967Proxy(
        address(new GatewayEVM()),
        abi.encodeWithSelector(GatewayEVM.initialize.selector, (echidnaCaller, address(0x123)))
    ));
    GatewayEVM testGateway = GatewayEVM(proxy);

    constructor() ERC20CustodyNew(address(testGateway)) {
        testERC20 = new TestERC20("test", "TEST");
        testGateway.setCustody(address(this));
    }

    function testWithdrawAndCall(address to, uint256 amount, bytes calldata data) public {
        // mint more than amount
        testERC20.mint(address(this), amount + 5);
        // transfer overhead to gateway
        testERC20.transfer(address(testGateway), 5);

        withdrawAndCall(address(testERC20), to, amount, data);

        // Assertion to ensure no ERC20 tokens are left in the gateway contract after execution
        assert(testERC20.balanceOf(address(gateway)) == 0);
    }
}