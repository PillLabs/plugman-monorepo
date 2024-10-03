import { Contract } from "ethers";
import { ethers, network } from "hardhat";
import { getAddress, isProtocolNetworkName } from "lib";

import { deployZetaConnectorEth, deployZetaConnectorNonEth, isEthNetworkName } from "../../../lib/contracts.helpers";

export async function deployZetaConnector(zetaTokenAddress: string) {
  if (!isProtocolNetworkName(network.name)) {
    throw new Error(`network.name: ${network.name} isn't supported.`);
  }
  const accounts = await ethers.getSigners();
  const [signer] = accounts;
  const initialBalance = await signer.getBalance();

  const tssAddress = getAddress("tss", network.name);
  const tssUpdaterAddress = getAddress("tssUpdater", network.name);

  const constructorArgs = [zetaTokenAddress, tssAddress, tssUpdaterAddress, tssUpdaterAddress];
  let contract: Contract;
  console.log(`Deploying ZetaConnector to ${network.name}`);

  if (isEthNetworkName(network.name)) {
    contract = await deployZetaConnectorEth({
      args: constructorArgs,
    });
  } else {
    contract = await deployZetaConnectorNonEth({
      args: constructorArgs,
    });
  }

  const finalBalance = await signer.getBalance();
  console.log("Deployed ZetaConnector. Address:", contract.address);
  console.log("Constructor Args", constructorArgs);
  console.log("ETH spent:", initialBalance.sub(finalBalance).toString());
  return contract.address;
}
