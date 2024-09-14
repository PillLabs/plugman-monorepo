// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "forge-std/Script.sol";
import "../src/MintVerifier.sol";

contract DeployMintVerifier is Script {
    function run() external {
        string memory name = "Plugman";
        address metadataCreator = vm.envAddress("METADATA_CREATOR");
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY_WITH_0x");
        vm.startBroadcast(deployerPrivateKey);
        MintVerifier verifier;
        verifier = new MintVerifier(name, metadataCreator);
        address verifierAddress = address(verifier);
        console.log("Verifier contract is deployed at", verifierAddress);
        vm.stopBroadcast();
    }
}