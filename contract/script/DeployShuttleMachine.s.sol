// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "forge-std/Script.sol";
import "../src/ShuttleMachine.sol";


contract DeployShuttleMachine is Script {
    function run() external {
        address zetaSystemContract = vm.envAddress("ZETA_SYSTEM_CONTRACT");
        address treasuryAddress = vm.envAddress("TREASURY_ADDRESS");
        uint256 permitChainId =  vm.envUint("PERMIT_CHAIN_ID");
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY_WITH_0x");
        address plugmanAddress = vm.envAddress("PLUGMAN_ADDRESS");
        uint256 crossingWlPrice = vm.envUint("CROSSING_WL_PRICE");
        uint256 crossingPublicSalePrice = vm.envUint("CROSSING_PUBLIC_SALE_PRICE");
        vm.startBroadcast(deployerPrivateKey);
        ShuttleMachine shuttleMachine;
        shuttleMachine = new ShuttleMachine(
            zetaSystemContract, plugmanAddress, treasuryAddress, permitChainId, crossingWlPrice, crossingPublicSalePrice);
        address shuttleMachineAddress = address(shuttleMachine);
        console.log("Shuttle machine contract is deployed at", shuttleMachineAddress);
        vm.stopBroadcast();
    }
}