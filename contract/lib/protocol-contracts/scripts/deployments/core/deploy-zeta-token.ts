import { Contract } from "ethers";
import { ethers, network } from "hardhat";
import { getAddress, isProtocolNetworkName } from "lib";

import { ZETA_INITIAL_SUPPLY } from "../../../lib/contracts.constants";
import { deployZetaEth, deployZetaNonEth, isEthNetworkName } from "../../../lib/contracts.helpers";

export async function deployZetaToken() {
  if (!isProtocolNetworkName(network.name)) {
    throw new Error(`network.name: ${network.name} isn't supported.`);
  }

  const accounts = await ethers.getSigners();
  const [signer] = accounts;
  const initialBalance = await signer.getBalance();

  const DEPLOYER_ADDRESS = process.env.DEPLOYER_ADDRESS || signer.address;

  const tssAddress = getAddress("tss", network.name);
  const tssUpdaterAddress = getAddress("tssUpdater", network.name);

  let contract: Contract;
  let constructorArgs: any;
  console.log(`Deploying ZetaToken to ${network.name}`);

  if (isEthNetworkName(network.name)) {
    constructorArgs = [DEPLOYER_ADDRESS, ZETA_INITIAL_SUPPLY];
    contract = await deployZetaEth({
      args: constructorArgs,
    });
  } else {
    constructorArgs = [tssAddress, tssUpdaterAddress];
    contract = await deployZetaNonEth({
      args: constructorArgs,
    });
  }

  const finalBalance = await signer.getBalance();
  console.log("Deployed ZetaToken. Address:", contract.address);
  console.log("Constructor Args", constructorArgs);
  console.log("ETH spent:", initialBalance.sub(finalBalance).toString());
  return contract.address;
}
