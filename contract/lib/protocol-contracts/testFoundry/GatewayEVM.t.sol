// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "forge-std/Test.sol";
import "forge-std/Vm.sol";

import "contracts/prototypes/evm/GatewayEVM.sol";
import "contracts/prototypes/evm/ReceiverEVM.sol";
import "contracts/prototypes/evm/ERC20CustodyNew.sol";
import "contracts/prototypes/evm/ZetaConnectorNonNative.sol";
import "contracts/prototypes/evm/TestERC20.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {Upgrades} from "openzeppelin-foundry-upgrades/LegacyUpgrades.sol";

import "contracts/prototypes/evm/IGatewayEVM.sol";
import "contracts/prototypes/evm/IERC20CustodyNew.sol";
import "contracts/prototypes/evm/IReceiverEVM.sol";

contract GatewayEVMTest is Test, IGatewayEVMErrors, IGatewayEVMEvents, IReceiverEVMEvents, IERC20CustodyNewEvents {
    using SafeERC20 for IERC20;

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
            abi.encodeWithSelector(GatewayEVM.initialize.selector, tssAddress, address(zeta))
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

    function testForwardCallToReceiveNonPayable() public {
        string[] memory str = new string[](1);
        str[0] = "Hello, Foundry!";
        uint256[] memory num = new uint256[](1);
        num[0] = 42;
        bool flag = true;

        bytes memory data = abi.encodeWithSignature("receiveNonPayable(string[],uint256[],bool)", str, num, flag);
        
        vm.expectCall(address(receiver), 0, data);
        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedNonPayable(address(gateway), str, num, flag);
        vm.expectEmit(true, true, true, true, address(gateway));
        emit Executed(address(receiver), 0, data);
        vm.prank(tssAddress);
        gateway.execute(address(receiver), data);
    }

    function testForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS() public {
        string[] memory str = new string[](1);
        str[0] = "Hello, Foundry!";
        uint256[] memory num = new uint256[](1);
        num[0] = 42;
        bool flag = true;
        bytes memory data = abi.encodeWithSignature("receiveNonPayable(string[],uint256[],bool)", str, num, flag);
        
        vm.prank(owner);
        vm.expectRevert(InvalidSender.selector);
        gateway.execute(address(receiver), data);
    }

    function testForwardCallToReceivePayable() public {
        string memory str = "Hello, Foundry!";
        uint256 num = 42;
        bool flag = true;
        uint256 value = 1 ether;
        assertEq(0, address(receiver).balance);

        bytes memory data = abi.encodeWithSignature("receivePayable(string,uint256,bool)", str, num, flag);
        
        vm.expectCall(address(receiver), 1 ether, data);
        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedPayable(address(gateway), value, str, num, flag);
        vm.expectEmit(true, true, true, true, address(gateway));
        emit Executed(address(receiver), 1 ether, data);
        vm.prank(tssAddress);
        gateway.execute{value: value}(address(receiver), data);

        assertEq(value, address(receiver).balance);
    }

    function testForwardCallToReceiveNoParams() public {
        bytes memory data = abi.encodeWithSignature("receiveNoParams()");
        
        vm.expectCall(address(receiver), 0, data);
        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedNoParams(address(gateway));
        vm.expectEmit(true, true, true, true, address(gateway));
        emit Executed(address(receiver), 0, data);
        vm.prank(tssAddress);
        gateway.execute(address(receiver), data);
    }

    function testExecuteWithERC20FailsIfNotCustoryOrConnector() public {
        uint256 amount = 100000;
        bytes memory data = abi.encodeWithSignature("receiveERC20(uint256,address,address)", amount, address(token), destination);

        vm.prank(owner);
        vm.expectRevert(InvalidSender.selector);
        gateway.executeWithERC20(address(token), destination, amount, data);
    }

    function testRevertWithERC20FailsIfNotCustoryOrConnector() public {
        uint256 amount = 100000;
        bytes memory data = abi.encodeWithSignature("receiveERC20(uint256,address,address)", amount, address(token), destination);

        vm.prank(owner);
        vm.expectRevert(InvalidSender.selector);
        gateway.revertWithERC20(address(token), destination, amount, data);
    }

    function testForwardCallToReceiveERC20ThroughCustody() public {
        uint256 amount = 100000;
        bytes memory data = abi.encodeWithSignature("receiveERC20(uint256,address,address)", amount, address(token), destination);
        uint256 balanceBefore = token.balanceOf(destination);
        assertEq(balanceBefore, 0);
        uint256 balanceBeforeCustody = token.balanceOf(address(custody));

        bytes memory transferData = abi.encodeWithSignature("transfer(address,uint256)", address(gateway), amount);
        vm.expectCall(address(token), 0, transferData);
        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedERC20(address(gateway), amount, address(token), destination);
        vm.expectEmit(true, true, true, true, address(custody));
        emit WithdrawAndCall(address(token), address(receiver), amount, data);
        vm.prank(tssAddress);
        custody.withdrawAndCall(address(token), address(receiver), amount, data);

        // Verify that the tokens were transferred to the destination address
        uint256 balanceAfter = token.balanceOf(destination);
        assertEq(balanceAfter, amount);

        // Verify that the remaining tokens were refunded to the Custody contract
        uint256 balanceAfterCustody = token.balanceOf(address(custody));
        assertEq(balanceAfterCustody, balanceBeforeCustody - amount);

        // Verify that the approval was reset
        uint256 allowance = token.allowance(address(gateway), address(receiver));
        assertEq(allowance, 0);

        // Verify that gateway doesn't hold any tokens
        uint256 balanceGateway = token.balanceOf(address(gateway));
        assertEq(balanceGateway, 0);
    }

    function testForwardCallToReceiveERC20ThroughCustodyFailsIfSenderIsNotTSS() public {
        uint256 amount = 100000;
        bytes memory data = abi.encodeWithSignature("receiveERC20(uint256,address,address)", amount, address(token), destination);

        vm.prank(owner);
        vm.expectRevert(InvalidSender.selector);
        custody.withdrawAndCall(address(token), address(receiver), amount, data);
    }

    function testForwardCallToReceiveERC20ThroughCustodyFailsIfAmountIs0() public {
        uint256 amount = 0;
        bytes memory data = abi.encodeWithSignature("receiveERC20(uint256,address,address)", amount, address(token), destination);
        
        vm.prank(tssAddress);
        vm.expectRevert(InsufficientERC20Amount.selector);
        custody.withdrawAndCall(address(token), address(receiver), amount, data);
    }

    function testForwardCallToReceiveERC20PartialThroughCustody() public {
        uint256 amount = 100000;
        bytes memory data = abi.encodeWithSignature("receiveERC20Partial(uint256,address,address)", amount, address(token), destination);
        uint256 balanceBefore = token.balanceOf(destination);
        assertEq(balanceBefore, 0);
        uint256 balanceBeforeCustody = token.balanceOf(address(custody));

        bytes memory transferData = abi.encodeWithSignature("transfer(address,uint256)", address(gateway), amount);
        vm.expectCall(address(token), 0, transferData);
        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedERC20(address(gateway), amount / 2, address(token), destination);
        vm.expectEmit(true, true, true, true, address(custody));
        emit WithdrawAndCall(address(token), address(receiver), amount, data);
        vm.prank(tssAddress);
        custody.withdrawAndCall(address(token), address(receiver), amount, data);

        // Verify that the tokens were transferred to the destination address
        uint256 balanceAfter = token.balanceOf(destination);
        assertEq(balanceAfter, amount / 2);

        // Verify that the remaining tokens were refunded to the Custody contract
        uint256 balanceAfterCustody = token.balanceOf(address(custody));
        assertEq(balanceAfterCustody, balanceBeforeCustody - amount / 2);

        // Verify that the approval was reset
        uint256 allowance = token.allowance(address(gateway), address(receiver));
        assertEq(allowance, 0);

        // Verify that gateway doesn't hold any tokens
        uint256 balanceGateway = token.balanceOf(address(gateway));
        assertEq(balanceGateway, 0);
    }

    function testForwardCallToReceiveERC20PartialThroughCustodyFailsIfSenderIsNotTSS() public {
        uint256 amount = 100000;
        bytes memory data = abi.encodeWithSignature("receiveERC20Partial(uint256,address,address)", amount, address(token), destination);
       
        vm.prank(owner);
        vm.expectRevert(InvalidSender.selector);
        custody.withdrawAndCall(address(token), address(receiver), amount, data);
    }

    function testForwardCallToReceiveERC20PartialThroughCustodyFailsIfAmountIs0() public {
        uint256 amount = 0;
        bytes memory data = abi.encodeWithSignature("receiveERC20Partial(uint256,address,address)", amount, address(token), destination);
        
        vm.prank(tssAddress);
        vm.expectRevert(InsufficientERC20Amount.selector);
        custody.withdrawAndCall(address(token), address(receiver), amount, data);
    }

    function testForwardCallToReceiveNoParamsThroughCustody() public {
        uint256 amount = 100000;
        bytes memory data = abi.encodeWithSignature("receiveNoParams()");
        uint256 balanceBefore = token.balanceOf(destination);
        assertEq(balanceBefore, 0);
        uint256 balanceBeforeCustody = token.balanceOf(address(custody));

        bytes memory transferData = abi.encodeWithSignature("transfer(address,uint256)", address(gateway), amount);
        vm.expectCall(address(token), 0, transferData);
        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedNoParams(address(gateway));
        vm.expectEmit(true, true, true, true, address(custody));
        emit WithdrawAndCall(address(token), address(receiver), amount, data);
        vm.prank(tssAddress);
        custody.withdrawAndCall(address(token), address(receiver), amount, data);

        // Verify that the tokens were not transferred to the destination address
        uint256 balanceAfter = token.balanceOf(destination);
        assertEq(balanceAfter, 0);

        // Verify that the remaining tokens were refunded to the Custody contract
        uint256 balanceAfterCustody = token.balanceOf(address(custody));
        assertEq(balanceAfterCustody, balanceBeforeCustody);

        // Verify that the approval was reset
        uint256 allowance = token.allowance(address(gateway), address(receiver));
        assertEq(allowance, 0);

        // Verify that gateway doesn't hold any tokens
        uint256 balanceGateway = token.balanceOf(address(gateway));
        assertEq(balanceGateway, 0);
    }

    function testWithdrawThroughCustody() public {
        uint256 amount = 100000;
        uint256 balanceBefore = token.balanceOf(destination);
        assertEq(balanceBefore, 0);
        uint256 balanceBeforeCustody = token.balanceOf(address(custody));

        bytes memory transferData = abi.encodeWithSignature("transfer(address,uint256)", address(destination), amount);
        vm.expectCall(address(token), 0, transferData);
        vm.expectEmit(true, true, true, true, address(custody));
        emit Withdraw(address(token), destination, amount);
        vm.prank(tssAddress);
        custody.withdraw(address(token), destination, amount);

        // Verify that the tokens were transferred to the destination address
        uint256 balanceAfter = token.balanceOf(destination);
        assertEq(balanceAfter, amount);

        // Verify that the tokens were substracted from custody
        uint256 balanceAfterCustody = token.balanceOf(address(custody));
        assertEq(balanceAfterCustody, balanceBeforeCustody - amount);

        // Verify that gateway doesn't hold any tokens
        uint256 balanceGateway = token.balanceOf(address(gateway));
        assertEq(balanceGateway, 0);
    }

    function testWithdrawThroughCustodyFailsIfSenderIsNotTSS() public {
        uint256 amount = 100000;

        vm.prank(owner);
        vm.expectRevert(InvalidSender.selector);
        custody.withdraw(address(token), destination, amount);
    }

    function testWithdrawAndRevertThroughCustody() public {
        uint256 amount = 100000;
        bytes memory data = abi.encodePacked("hello");
        uint256 balanceBefore = token.balanceOf(address(receiver));
        assertEq(balanceBefore, 0);
        uint256 balanceBeforeCustody = token.balanceOf(address(custody));

        bytes memory transferData = abi.encodeWithSignature("transfer(address,uint256)", address(gateway), amount);
        vm.expectCall(address(token), 0, transferData);
        // Verify that onRevert callback was called
        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedRevert(address(gateway), data);
        vm.expectEmit(true, true, true, true, address(gateway));
        emit RevertedWithERC20(address(token), address(receiver), amount, data);
        vm.expectEmit(true, true, true, true, address(custody));
        emit WithdrawAndRevert(address(token), address(receiver), amount, data);
        vm.prank(tssAddress);
        custody.withdrawAndRevert(address(token), address(receiver), amount, data);

        // Verify that the tokens were transferred to the receiver address
        uint256 balanceAfter = token.balanceOf(address(receiver));
        assertEq(balanceAfter, amount);

        // Verify that zeta connector doesn't get more tokens
        uint256 balanceAfterCustody = token.balanceOf(address(custody));
        assertEq(balanceAfterCustody, balanceBeforeCustody - amount);

        // Verify that the approval was reset
        uint256 allowance = token.allowance(address(gateway), address(receiver));
        assertEq(allowance, 0);

        // Verify that gateway doesn't hold any tokens
        uint256 balanceGateway = token.balanceOf(address(gateway));
        assertEq(balanceGateway, 0);
    }

    function testWithdrawAndRevertThroughCustodyFailsIfSenderIsNotTSS() public {
        uint256 amount = 100000;
        bytes memory data = abi.encodePacked("hello");

        vm.prank(owner);
        vm.expectRevert(InvalidSender.selector);
        custody.withdrawAndRevert(address(token), address(receiver), amount, data);
    }

    function testWithdrawAndRevertThroughCustodyFailsIfAmountIs0() public {
        uint256 amount = 0;
        bytes memory data = abi.encodePacked("hello");

        vm.prank(tssAddress);
        vm.expectRevert(InsufficientERC20Amount.selector);
        custody.withdrawAndRevert(address(token), address(receiver), amount, data);
    }

    function testExecuteRevert() public {
        uint256 value = 1 ether;
        bytes memory data = abi.encodePacked("hello");
        uint256 balanceBefore = address(receiver).balance;
        assertEq(balanceBefore, 0);

        // Verify that onRevert callback was called
        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedRevert(address(gateway), data);
        vm.expectEmit(true, true, true, true, address(gateway));
        emit Reverted(address(receiver), 1 ether, data);
        vm.prank(tssAddress);
        gateway.executeRevert{value: value}(address(receiver), data);

        // Verify that the tokens were transferred to the receiver address
        uint256 balanceAfter = address(receiver).balance;
        assertEq(balanceAfter, 1 ether);
    }

    function testExecuteRevertFailsIfSenderIsNotTSS() public {
        uint256 value = 1 ether;
        bytes memory data = abi.encodePacked("hello");

        vm.prank(owner);
        vm.expectRevert(InvalidSender.selector);
        gateway.executeRevert{value: value}(address(receiver), data);
    }
}

contract GatewayEVMInboundTest is Test, IGatewayEVMErrors, IGatewayEVMEvents, IReceiverEVMEvents {
    using SafeERC20 for IERC20;

    GatewayEVM gateway;
    ERC20CustodyNew custody;
    ZetaConnectorNonNative zetaConnector;
    TestERC20 token;
    TestERC20 zeta;
    address owner;
    address destination;
    address tssAddress;

    uint256 ownerAmount = 1000000;

    function setUp() public {
        owner = address(this);
        destination = address(0x1234);
        tssAddress = address(0x5678);

        token = new TestERC20("test", "TTK");
        zeta = new TestERC20("zeta", "ZETA");
        address proxy = address(new ERC1967Proxy(
            address(new GatewayEVM()),
            abi.encodeWithSelector(GatewayEVM.initialize.selector, tssAddress, address(zeta))
        ));
        gateway = GatewayEVM(proxy);
        custody = new ERC20CustodyNew(address(gateway), tssAddress);
        zetaConnector = new ZetaConnectorNonNative(address(gateway), address(zeta), tssAddress);

        vm.deal(tssAddress, 1 ether);

        vm.startPrank(tssAddress);
        gateway.setCustody(address(custody));
        gateway.setConnector(address(zetaConnector));
        vm.stopPrank();

        token.mint(owner, ownerAmount);
    }

    function testDepositERC20ToCustody() public {
        uint256 amount = 100000;
        uint256 custodyBalanceBefore = token.balanceOf(address(custody));
        assertEq(0, custodyBalanceBefore);

        token.approve(address(gateway), amount);

        vm.expectEmit(true, true, true, true, address(gateway));
        emit Deposit(owner, destination, amount, address(token), "");
        gateway.deposit(destination, amount, address(token));

        uint256 custodyBalanceAfter = token.balanceOf(address(custody));
        assertEq(amount, custodyBalanceAfter);

        uint256 ownerAmountAfter = token.balanceOf(owner);
        assertEq(ownerAmount - amount, ownerAmountAfter);
    }

    function testFailDepositERC20ToCustodyIfAmountIs0() public {
        uint256 amount = 0;

        token.approve(address(gateway), amount);

        vm.expectRevert("InsufficientERC20Amount");
        gateway.deposit(destination, amount, address(token));
    }

    function testDepositEthToTss() public {
        uint256 amount = 100000;
        uint256 tssBalanceBefore = tssAddress.balance;

        vm.expectEmit(true, true, true, true, address(gateway));
        emit Deposit(owner, destination, amount, address(0), "");
        gateway.deposit{value: amount}(destination);

        uint256 tssBalanceAfter = tssAddress.balance;
        assertEq(tssBalanceBefore + amount, tssBalanceAfter);
    }

    function testFailDepositEthToTssIfAmountIs0() public {
        uint256 amount = 0;

        vm.expectRevert("InsufficientETHAmount");
        gateway.deposit{value: amount}(destination);
    }

    function testDepositERC20ToCustodyWithPayload() public {
        uint256 amount = 100000;
        uint256 custodyBalanceBefore = token.balanceOf(address(custody));
        assertEq(0, custodyBalanceBefore);

        bytes memory payload = abi.encodeWithSignature("hello(address)", destination);

        token.approve(address(gateway), amount);

        vm.expectEmit(true, true, true, true, address(gateway));
        emit Deposit(owner, destination, amount, address(token), payload);
        gateway.depositAndCall(destination, amount, address(token), payload);

        uint256 custodyBalanceAfter = token.balanceOf(address(custody));
        assertEq(amount, custodyBalanceAfter);

        uint256 ownerAmountAfter = token.balanceOf(owner);
        assertEq(ownerAmount - amount, ownerAmountAfter);
    }

    function testFailDepositERC20ToCustodyWithPayloadIfAmountIs0() public {
        uint256 amount = 0;

        bytes memory payload = abi.encodeWithSignature("hello(address)", destination);

        vm.expectRevert("InsufficientERC20Amount");
        gateway.depositAndCall(destination, amount, address(token), payload);
    }

    function testDepositEthToTssWithPayload() public {
        uint256 amount = 100000;
        uint256 tssBalanceBefore = tssAddress.balance;
        bytes memory payload = abi.encodeWithSignature("hello(address)", destination);

        vm.expectEmit(true, true, true, true, address(gateway));
        emit Deposit(owner, destination, amount, address(0), payload);
        gateway.depositAndCall{value: amount}(destination, payload);

        uint256 tssBalanceAfter = tssAddress.balance;
        assertEq(tssBalanceBefore + amount, tssBalanceAfter);
    }

    function testFailDepositEthToTssWithPayloadIfAmountIs0() public {
        uint256 amount = 0;
        bytes memory payload = abi.encodeWithSignature("hello(address)", destination);

        vm.expectRevert("InsufficientETHAmount");
        gateway.depositAndCall{value: amount}(destination, payload);
    }

    function testCallWithPayload() public {
        bytes memory payload = abi.encodeWithSignature("hello(address)", destination);

        vm.expectEmit(true, true, true, true, address(gateway));
        emit Call(owner, destination, payload);
        gateway.call(destination, payload);
    }
}