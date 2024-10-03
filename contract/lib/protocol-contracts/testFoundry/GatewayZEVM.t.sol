// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "forge-std/Test.sol";
import "forge-std/Vm.sol";

import "contracts/prototypes/zevm/GatewayZEVM.sol";
import "contracts/zevm/ZRC20New.sol";
import "contracts/zevm/SystemContract.sol";
import "contracts/zevm/interfaces/IZRC20.sol";
import "contracts/prototypes/zevm/TestZContract.sol";
import "contracts/prototypes/zevm/IGatewayZEVM.sol";
import "contracts/zevm/WZETA.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {Upgrades} from "openzeppelin-foundry-upgrades/LegacyUpgrades.sol";
import "forge-std/console.sol";

contract GatewayZEVMInboundTest is Test, IGatewayZEVMEvents, IGatewayZEVMErrors {
    address payable proxy;
    GatewayZEVM gateway;
    ZRC20New zrc20;
    WETH9 zetaToken;
    SystemContract systemContract;
    TestZContract testZContract;
    address owner;
    address addr1;
    address fungibleModule;

    function setUp() public {
        owner = address(this);
        addr1 = address(0x1234);

        zetaToken = new WETH9();

        proxy = payable(address(new ERC1967Proxy(
            address(new GatewayZEVM()),
            abi.encodeWithSelector(GatewayZEVM.initialize.selector, address(zetaToken))
        )));
        gateway = GatewayZEVM(proxy);
        fungibleModule = gateway.FUNGIBLE_MODULE_ADDRESS();
        testZContract = new TestZContract();

        vm.startPrank(fungibleModule);
        systemContract = new SystemContract(address(0), address(0), address(0));
        zrc20 = new ZRC20New("TOKEN", "TKN", 18, 1, CoinType.Gas, 0, address(systemContract), address(gateway));
        systemContract.setGasCoinZRC20(1, address(zrc20));
        systemContract.setGasPrice(1, 1);
        vm.deal(fungibleModule, 1000000000);
        zetaToken.deposit{value: 10}();
        zetaToken.approve(address(gateway), 10);
        zrc20.deposit(owner, 100000);
        vm.stopPrank();

        vm.startPrank(owner);
        zrc20.approve(address(gateway), 100000);
        zetaToken.deposit{value: 10}();
        zetaToken.approve(address(gateway), 10);
        vm.stopPrank();
    }

    function testWithdrawZRC20() public {
        uint256 amount = 1;
        uint256 ownerBalanceBefore = zrc20.balanceOf(owner);

        vm.expectEmit(true, true, true, true, address(gateway));
        emit Withdrawal(owner, address(zrc20), abi.encodePacked(addr1), amount, 0, zrc20.PROTOCOL_FLAT_FEE(), "");
        gateway.withdraw(abi.encodePacked(addr1), 1, address(zrc20));

        uint256 ownerBalanceAfter = zrc20.balanceOf(owner);
        assertEq(ownerBalanceBefore - amount, ownerBalanceAfter);
    }

    function testWithdrawZRC20FailsIfNoAllowance() public {
        uint256 amount = 1;
        uint256 ownerBalanceBefore = zrc20.balanceOf(owner);
        // Remove allowance for gateway
        vm.prank(owner);
        zrc20.approve(address(gateway), 0);

        vm.expectRevert();
        gateway.withdraw(abi.encodePacked(addr1), amount, address(zrc20));

        // Check that balance didn't change
        uint256 ownerBalanceAfter = zrc20.balanceOf(owner);
        assertEq(ownerBalanceBefore, ownerBalanceAfter);
    }

    function testWithdrawZRC20WithMessageFailsIfNoAllowance() public {
        uint256 amount = 1;
        uint256 ownerBalanceBefore = zrc20.balanceOf(owner);
        // Remove allowance for gateway
        vm.prank(owner);
        zrc20.approve(address(gateway), 0);

        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        vm.expectRevert();
        gateway.withdrawAndCall(abi.encodePacked(addr1), amount, address(zrc20), message);

        // Check that balance didn't change
        uint256 ownerBalanceAfter = zrc20.balanceOf(owner);
        assertEq(ownerBalanceBefore, ownerBalanceAfter);
    }

    function testWithdrawZRC20WithMessage() public {
        uint256 amount = 1;
        uint256 ownerBalanceBefore = zrc20.balanceOf(owner);

        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        vm.expectEmit(true, true, true, true, address(gateway));
        emit Withdrawal(owner, address(zrc20), abi.encodePacked(addr1), amount, 0, zrc20.PROTOCOL_FLAT_FEE(), message);
        gateway.withdrawAndCall(abi.encodePacked(addr1), amount, address(zrc20), message);

        uint256 ownerBalanceAfter = zrc20.balanceOf(owner);
        assertEq(ownerBalanceBefore - amount, ownerBalanceAfter);
    }

    function testWithdrawZETA() public {
        uint256 amount = 1;
        uint256 ownerBalanceBefore = zetaToken.balanceOf(owner);
        uint256 gatewayBalanceBefore = zetaToken.balanceOf(address(gateway));
        uint256 fungibleModuleBalanceBefore = fungibleModule.balance;

        vm.expectEmit(true, true, true, true, address(gateway));
        emit Withdrawal(owner, address(zetaToken), abi.encodePacked(fungibleModule), amount, 0, 0, "");
        gateway.withdraw(amount);

        uint256 ownerBalanceAfter = zetaToken.balanceOf(owner);
        assertEq(ownerBalanceBefore - 1, ownerBalanceAfter);

        uint256 gatewayBalanceAfter = zetaToken.balanceOf(address(gateway));
        assertEq(gatewayBalanceBefore, gatewayBalanceAfter);
        
        // Verify amount is transfered to fungible module
        assertEq(fungibleModuleBalanceBefore + 1, fungibleModule.balance);
    }

    function testWithdrawZETAFailsIfNoAllowance() public {
        uint256 amount = 1;
        uint256 ownerBalanceBefore = zetaToken.balanceOf(owner);
        uint256 gatewayBalanceBefore = zetaToken.balanceOf(address(gateway));
        uint256 fungibleModuleBalanceBefore = fungibleModule.balance;
        // Remove allowance for gateway
        vm.prank(owner);
        zetaToken.approve(address(gateway), 0);

        vm.expectRevert();
        gateway.withdraw(amount);

        // Verify balances not changed
        uint256 ownerBalanceAfter = zetaToken.balanceOf(owner);
        assertEq(ownerBalanceBefore, ownerBalanceAfter);

        uint256 gatewayBalanceAfter = zetaToken.balanceOf(address(gateway));
        assertEq(gatewayBalanceBefore, gatewayBalanceAfter);
        
        assertEq(fungibleModuleBalanceBefore, fungibleModule.balance);
    }

    function testWithdrawZETAWithMessage() public {
        uint256 amount = 1;
        uint256 ownerBalanceBefore = zetaToken.balanceOf(owner);
        uint256 gatewayBalanceBefore = zetaToken.balanceOf(address(gateway));
        uint256 fungibleModuleBalanceBefore = fungibleModule.balance;
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);

        vm.expectEmit(true, true, true, true, address(gateway));
        emit Withdrawal(owner, address(zetaToken), abi.encodePacked(fungibleModule), amount, 0, 0, message);
        gateway.withdrawAndCall(amount, message);

        uint256 ownerBalanceAfter = zetaToken.balanceOf(owner);
        assertEq(ownerBalanceBefore - 1, ownerBalanceAfter);

        uint256 gatewayBalanceAfter = zetaToken.balanceOf(address(gateway));
        assertEq(gatewayBalanceBefore, gatewayBalanceAfter);
        
        // Verify amount is transfered to fungible module
        assertEq(fungibleModuleBalanceBefore + 1, fungibleModule.balance);
    }

    function testWithdrawZETAWithMessageFailsIfNoAllowance() public {
        uint256 amount = 1;
        uint256 ownerBalanceBefore = zetaToken.balanceOf(owner);
        uint256 gatewayBalanceBefore = zetaToken.balanceOf(address(gateway));
        uint256 fungibleModuleBalanceBefore = fungibleModule.balance;
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        // Remove allowance for gateway
        vm.prank(owner);
        zetaToken.approve(address(gateway), 0);

        vm.expectRevert();
        gateway.withdrawAndCall(amount, message);

        // Verify balances not changed
        uint256 ownerBalanceAfter = zetaToken.balanceOf(owner);
        assertEq(ownerBalanceBefore, ownerBalanceAfter);

        uint256 gatewayBalanceAfter = zetaToken.balanceOf(address(gateway));
        assertEq(gatewayBalanceBefore, gatewayBalanceAfter);
        
        assertEq(fungibleModuleBalanceBefore, fungibleModule.balance);
    }

    function testCall() public {
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        vm.expectEmit(true, true, true, true, address(gateway));
        emit Call(owner, abi.encodePacked(addr1), message);
        gateway.call(abi.encodePacked(addr1), message);
    }
}

contract GatewayZEVMOutboundTest is Test, IGatewayZEVMEvents, IGatewayZEVMErrors {
    address payable proxy;
    GatewayZEVM gateway;
    ZRC20New zrc20;
    WETH9 zetaToken;
    SystemContract systemContract;
    TestZContract testZContract;
    address owner;
    address addr1;
    address fungibleModule;

    event ContextData(bytes origin, address sender, uint256 chainID, address msgSender, string message);
    event ContextDataRevert(bytes origin, address sender, uint256 chainID, address msgSender, string message);

    function setUp() public {
        owner = address(this);
        addr1 = address(0x1234);

        zetaToken = new WETH9();

        proxy = payable(address(new ERC1967Proxy(
            address(new GatewayZEVM()),
            abi.encodeWithSelector(GatewayZEVM.initialize.selector, address(zetaToken))
        )));
        gateway = GatewayZEVM(proxy);
        fungibleModule = gateway.FUNGIBLE_MODULE_ADDRESS();

        testZContract = new TestZContract();

        vm.startPrank(fungibleModule);
        systemContract = new SystemContract(address(0), address(0), address(0));
        zrc20 = new ZRC20New("TOKEN", "TKN", 18, 1, CoinType.Gas, 0, address(systemContract), address(gateway));
        systemContract.setGasCoinZRC20(1, address(zrc20));
        systemContract.setGasPrice(1, 1);
        vm.deal(fungibleModule, 1000000000);
        zetaToken.deposit{value: 10}();
        zetaToken.approve(address(gateway), 10);
        zrc20.deposit(owner, 100000);
        vm.stopPrank();

        vm.startPrank(owner);
        zrc20.approve(address(gateway), 100000);
        zetaToken.deposit{value: 10}();
        zetaToken.approve(address(gateway), 10);
        vm.stopPrank();
    }

    function testDeposit() public {
        uint256 amount = 1;
        uint256 balanceBefore = zrc20.balanceOf(addr1);
        assertEq(0, balanceBefore);

        vm.prank(fungibleModule);
        gateway.deposit(address(zrc20), amount, addr1);

        uint256 balanceAfter = zrc20.balanceOf(addr1);
        assertEq(amount, balanceAfter);
    }

    function testDepositFailsIfSenderNotFungibleModule() public {
        uint256 amount = 1;
        uint256 balanceBefore = zrc20.balanceOf(addr1);
        assertEq(0, balanceBefore);

        vm.prank(owner);
        vm.expectRevert(CallerIsNotFungibleModule.selector);
        gateway.deposit(address(zrc20), amount, addr1);

        uint256 balanceAfter = zrc20.balanceOf(addr1);
        assertEq(0, balanceAfter);
    }

    function testDepositFailsIfTargetIsGateway() public {
        uint256 amount = 1;

        vm.prank(fungibleModule);
        vm.expectRevert(InvalidTarget.selector);
        gateway.deposit(address(zrc20), amount, address(gateway));
    }

    function testDepositFailsIfTargetIsFungibleModule() public {
        uint256 amount = 1;
        vm.prank(fungibleModule);
        vm.expectRevert(InvalidTarget.selector);
        gateway.deposit(address(zrc20), amount, fungibleModule);
    }

    function testExecuteZContract() public {
        bytes memory message = abi.encode("hello");
        zContext memory context = zContext({
            origin: abi.encodePacked(address(gateway)),
            sender: fungibleModule,
            chainID: 1
        });

        vm.expectEmit(true, true, true, true, address(testZContract));
        emit ContextData(abi.encodePacked(gateway), fungibleModule, 1, address(gateway), "hello");
        vm.prank(fungibleModule);
        gateway.execute(context, address(zrc20), 1, address(testZContract), message);
    }

    function testExecuteZContractFailsIfSenderIsNotFungibleModule() public {
        bytes memory message = abi.encode("hello");
        zContext memory context = zContext({
            origin: abi.encodePacked(address(gateway)),
            sender: fungibleModule,
            chainID: 1
        });

        vm.expectRevert(CallerIsNotFungibleModule.selector);
        vm.prank(owner);
        gateway.execute(context, address(zrc20), 1, address(testZContract), message);
    }

    function testExecuteRevertZContract() public {
        bytes memory message = abi.encode("hello");
        revertContext memory context = revertContext({
            origin: abi.encodePacked(address(gateway)),
            sender: fungibleModule,
            chainID: 1
        });

        vm.expectEmit(true, true, true, true, address(testZContract));
        emit ContextDataRevert(abi.encodePacked(gateway), fungibleModule, 1, address(gateway), "hello");
        vm.prank(fungibleModule);
        gateway.executeRevert(context, address(zrc20), 1, address(testZContract), message);
    }

     function testExecuteRevertZContractIfSenderIsNotFungibleModule() public {
        bytes memory message = abi.encode("hello");
        revertContext memory context = revertContext({
            origin: abi.encodePacked(address(gateway)),
            sender: fungibleModule,
            chainID: 1
        });

        vm.expectRevert(CallerIsNotFungibleModule.selector);
        vm.prank(owner);
        gateway.executeRevert(context, address(zrc20), 1, address(testZContract), message);
    }
   
    function testDepositZRC20AndCallZContract() public {
        uint256 balanceBefore = zrc20.balanceOf(address(testZContract));
        assertEq(0, balanceBefore);

        bytes memory message = abi.encode("hello");
        zContext memory context = zContext({
            origin: abi.encodePacked(address(gateway)),
            sender: fungibleModule,
            chainID: 1
        });
    
        vm.expectEmit(true, true, true, true, address(testZContract));
        emit ContextData(abi.encodePacked(gateway), fungibleModule, 1, address(gateway), "hello");
        vm.prank(fungibleModule);
        gateway.depositAndCall(context, address(zrc20), 1, address(testZContract), message);

        uint256 balanceAfter = zrc20.balanceOf(address(testZContract));
        assertEq(1, balanceAfter);
    }

    function testDepositZRC20AndCallZContractFailsIfSenderIsNotFungibleModule() public {
        bytes memory message = abi.encode("hello");
        zContext memory context = zContext({
            origin: abi.encodePacked(address(gateway)),
            sender: fungibleModule,
            chainID: 1
        });
    
        vm.expectRevert(CallerIsNotFungibleModule.selector);
        vm.prank(owner);
        gateway.depositAndCall(context, address(zrc20), 1, address(testZContract), message);
    }

    function testDepositZRC20AndCallZContractIfTargetIsFungibleModule() public {
        bytes memory message = abi.encode("hello");
        zContext memory context = zContext({
            origin: abi.encodePacked(address(gateway)),
            sender: fungibleModule,
            chainID: 1
        });
    
        vm.expectRevert(InvalidTarget.selector);
        vm.prank(fungibleModule);
        gateway.depositAndCall(context, address(zrc20), 1, fungibleModule, message);
    }

    function testDepositZRC20AndCallZContractIfTargetIsGateway() public {
        bytes memory message = abi.encode("hello");
        zContext memory context = zContext({
            origin: abi.encodePacked(address(gateway)),
            sender: fungibleModule,
            chainID: 1
        });
    
        vm.expectRevert(InvalidTarget.selector);
        vm.prank(fungibleModule);
        gateway.depositAndCall(context, address(zrc20), 1, address(gateway), message);
    }

    function testDepositAndRevertZRC20AndCallZContract() public {
        uint256 balanceBefore = zrc20.balanceOf(address(testZContract));
        assertEq(0, balanceBefore);

        bytes memory message = abi.encode("hello");
        revertContext memory context = revertContext({
            origin: abi.encodePacked(address(gateway)),
            sender: fungibleModule,
            chainID: 1
        });
    
        vm.expectEmit(true, true, true, true, address(testZContract));
        emit ContextDataRevert(abi.encodePacked(gateway), fungibleModule, 1, address(gateway), "hello");
        vm.prank(fungibleModule);
        gateway.depositAndRevert(context, address(zrc20), 1, address(testZContract), message);

        uint256 balanceAfter = zrc20.balanceOf(address(testZContract));
        assertEq(1, balanceAfter);
    }

    function testDepositAndRevertZRC20AndCallZContractFailsIfSenderIsNotFungibleModule() public {
        bytes memory message = abi.encode("hello");
        revertContext memory context = revertContext({
            origin: abi.encodePacked(address(gateway)),
            sender: fungibleModule,
            chainID: 1
        });
    
        vm.expectRevert(CallerIsNotFungibleModule.selector);
        vm.prank(owner);
        gateway.depositAndRevert(context, address(zrc20), 1, address(testZContract), message);
    }

    function testDepositAndRevertZRC20AndCallZContractFailsITargetIsFungibleModule() public {
        bytes memory message = abi.encode("hello");
        revertContext memory context = revertContext({
            origin: abi.encodePacked(address(gateway)),
            sender: fungibleModule,
            chainID: 1
        });
    
        vm.expectRevert(InvalidTarget.selector);
        vm.prank(fungibleModule);
        gateway.depositAndRevert(context, address(zrc20), 1, fungibleModule, message);
    }

     function testDepositAndRevertZRC20AndCallZContractFailsITargetIsGateway() public {
        bytes memory message = abi.encode("hello");
        revertContext memory context = revertContext({
            origin: abi.encodePacked(address(gateway)),
            sender: fungibleModule,
            chainID: 1
        });
    
        vm.expectRevert(InvalidTarget.selector);
        vm.prank(fungibleModule);
        gateway.depositAndRevert(context, address(zrc20), 1, address(gateway), message);
    }

    function testDepositZETAAndCallZContract() public {
        uint256 amount = 1;
        uint256 fungibleBalanceBefore = zetaToken.balanceOf(fungibleModule);
        uint256 gatewayBalanceBefore = zetaToken.balanceOf(address(gateway));
        uint256 destinationBalanceBefore = address(testZContract).balance;
        bytes memory message = abi.encode("hello");
        zContext memory context = zContext({
            origin: abi.encodePacked(address(gateway)),
            sender: fungibleModule,
            chainID: 1
        });

        vm.expectEmit(true, true, true, true, address(testZContract));
        emit ContextData(abi.encodePacked(gateway), fungibleModule, amount, address(gateway), "hello");
        vm.prank(fungibleModule);
        gateway.depositAndCall(context, amount, address(testZContract), message);

        uint256 fungibleBalanceAfter = zetaToken.balanceOf(fungibleModule);
        assertEq(fungibleBalanceBefore - amount, fungibleBalanceAfter);

        uint256 gatewayBalanceAfter = zetaToken.balanceOf(address(gateway));
        assertEq(gatewayBalanceBefore, gatewayBalanceAfter);
        
        // Verify amount is transfered to destination
        assertEq(destinationBalanceBefore + amount, address(testZContract).balance);
    }

    function testDepositZETAAndCallZContractFailsIfSenderIsNotFungibleModule() public {
        uint256 amount = 1;
        bytes memory message = abi.encode("hello");
        zContext memory context = zContext({
            origin: abi.encodePacked(address(gateway)),
            sender: fungibleModule,
            chainID: 1
        });

        vm.expectRevert(CallerIsNotFungibleModule.selector);
        vm.prank(owner);
        gateway.depositAndCall(context, amount, address(testZContract), message);
    }

    function testDepositZETAAndCallZContractFailsIfTargetIsFungibleModule() public {
        uint256 amount = 1;
        bytes memory message = abi.encode("hello");
        zContext memory context = zContext({
            origin: abi.encodePacked(address(gateway)),
            sender: fungibleModule,
            chainID: 1
        });

        vm.expectRevert(InvalidTarget.selector);
        vm.prank(fungibleModule);
        gateway.depositAndCall(context, amount, fungibleModule, message);
    }

     function testDepositZETAAndCallZContractFailsIfTargetIsGateway() public {
        uint256 amount = 1;
        bytes memory message = abi.encode("hello");
        zContext memory context = zContext({
            origin: abi.encodePacked(address(gateway)),
            sender: fungibleModule,
            chainID: 1
        });

        vm.expectRevert(InvalidTarget.selector);
        vm.prank(fungibleModule);
        gateway.depositAndCall(context, amount, address(gateway), message);
    }
}