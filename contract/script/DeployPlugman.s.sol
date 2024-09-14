// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "forge-std/Script.sol";
import "../src/Plugman.sol";

contract DeployPlugman is Script {
    function run() external {
        string memory name = "Plugman";
        string memory symbol = "PLG";
        address verifierAddress = vm.envAddress("VERIFIER_ADDRESS");
        uint256 batchMax = vm.envUint("BATCH_MAX");
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY_WITH_0x");
        vm.startBroadcast(deployerPrivateKey);
        Plugman plugman;
        plugman = new Plugman(name, symbol, verifierAddress, batchMax);
        address plugmanAddress = address(plugman);
        console.log("Plugman contract is deployed at", plugmanAddress);
        vm.stopBroadcast();
    }
}