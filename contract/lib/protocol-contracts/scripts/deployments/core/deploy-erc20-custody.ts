import { Contract } from "ethers";
import { ethers, network } from "hardhat";
import { getAddress, isProtocolNetworkName } from "lib";

import { ERC20_CUSTODY_ZETA_FEE, ERC20_CUSTODY_ZETA_MAX_FEE } from "../../../lib/contracts.constants";
import { deployERC20Custody as deployERC20CustodyHelper } from "../../../lib/contracts.helpers";

export async function deployERC20Custody(zetaTokenAddress: string) {
  if (!isProtocolNetworkName(network.name)) {
    throw new Error(`network.name: ${network.name} isn't supported.`);
  }
  const accounts = await ethers.getSigners();
  const [signer] = accounts;
  const initialBalance = await signer.getBalance();

  const tssAddress = getAddress("tss", network.name);
  const tssUpdaterAddress = getAddress("tssUpdater", network.name);

  const zetaFee = ERC20_CUSTODY_ZETA_FEE;
  const zetaMaxFee = ERC20_CUSTODY_ZETA_MAX_FEE;

  let contract: Contract;
  console.log(`Deploying ERC20Custody to ${network.name}`);

  const constructorArgs = [tssAddress, tssUpdaterAddress, zetaFee.toString(), zetaMaxFee.toString(), zetaTokenAddress];
  contract = await deployERC20CustodyHelper({
    args: constructorArgs,
  });

  const finalBalance = await signer.getBalance();
  console.log("Deployed ERC20Custody. Address:", contract.address);
  console.log("Constructor Args", constructorArgs);
  console.log("ETH spent:", initialBalance.sub(finalBalance).toString());

  return contract.address;
}
