// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "forge-std/Test.sol";
import "forge-std/Vm.sol";

import "contracts/prototypes/evm/GatewayEVM.sol";
import "contracts/prototypes/evm/GatewayEVMUpgradeTest.sol";
import "contracts/prototypes/evm/ReceiverEVM.sol";
import "contracts/prototypes/evm/ERC20CustodyNew.sol";
import "contracts/prototypes/evm/ZetaConnectorNonNative.sol";
import "contracts/prototypes/evm/TestERC20.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {Upgrades} from "openzeppelin-foundry-upgrades/LegacyUpgrades.sol";

import "contracts/prototypes/evm/IGatewayEVM.sol";
import "contracts/prototypes/evm/IReceiverEVM.sol";
contract GatewayEVMUUPSUpgradeTest is Test, IGatewayEVMErrors, IGatewayEVMEvents, IReceiverEVMEvents {
    using SafeERC20 for IERC20;
    event ExecutedV2(address indexed destination, uint256 value, bytes data);

    address proxy;
    GatewayEVM gateway;
    ReceiverEVM receiver;
    ERC20CustodyNew custody;
    ZetaConnectorNonNative zetaConnector;
    TestERC20 token;
    TestERC20 zeta;
    address owner;
    address destination;
    address tssAddress;

    function setUp() public {
        owner = address(this);
        destination = address(0x1234);
        tssAddress = address(0x5678);

        token = new TestERC20("test", "TTK");
        zeta = new TestERC20("zeta", "ZETA");

        proxy = address(new ERC1967Proxy(
            address(new GatewayEVM()),
            abi.encodeWithSelector(GatewayEVM.initialize.selector, tssAddress, zeta)
        ));
        gateway = GatewayEVM(proxy);

        custody = new ERC20CustodyNew(address(gateway), tssAddress);
        zetaConnector = new ZetaConnectorNonNative(address(gateway), address(zeta), tssAddress);
        receiver = new ReceiverEVM();

        vm.deal(tssAddress, 1 ether);

        vm.startPrank(tssAddress);
        gateway.setCustody(address(custody));
        gateway.setConnector(address(zetaConnector));
        vm.stopPrank();

        token.mint(owner, 1000000);
        token.transfer(address(custody), 500000);

        vm.deal(tssAddress, 1 ether);
    }

    function testUpgradeAndForwardCallToReceivePayable() public {
        address custodyBeforeUpgrade = gateway.custody();
        address tssBeforeUpgrade = gateway.tssAddress();

        string memory str = "Hello, Foundry!";
        uint256 num = 42;
        bool flag = true;
        uint256 value = 1 ether;

        vm.prank(owner);
        Upgrades.upgradeProxy(
            proxy,
            "GatewayEVMUpgradeTest.sol",
            "",
            owner
        );

        bytes memory data = abi.encodeWithSignature("receivePayable(string,uint256,bool)", str, num, flag);
        GatewayEVMUpgradeTest gatewayUpgradeTest = GatewayEVMUpgradeTest(proxy);

        vm.expectCall(address(receiver), value, data);
        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedPayable(address(gateway), value, str, num, flag);
        vm.expectEmit(true, true, true, true, address(gateway));
        emit ExecutedV2(address(receiver), value, data);
        vm.prank(tssAddress);
        gatewayUpgradeTest.execute{value: value}(address(receiver), data);

        assertEq(custodyBeforeUpgrade, gateway.custody());
        assertEq(tssBeforeUpgrade, gateway.tssAddress());
    }
}