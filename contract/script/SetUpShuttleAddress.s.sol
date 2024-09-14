// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "forge-std/Script.sol";
import "../src/interface/IPlugman.sol";

// RPC=https://zetachain-testnet-archive.allthatnode.com:8545/
// forge script script/Deploy.s.sol --fork-url "${RPC}" --broadcast


contract SetUpShuttle is Script {
    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY_WITH_0x");
        address plugmanAddress = vm.envAddress("PLUGMAN_ADDRESS");
        address shuttleMachineAddress = vm.envAddress("SHUTTLE_MACHINE_ADDRESS");
        vm.startBroadcast(deployerPrivateKey);
        PlugmanInterface(plugmanAddress).setShuttleMachine(shuttleMachineAddress);
        console.log("Set shuttle machine contract address on plugman");
        vm.stopBroadcast();
    }
}