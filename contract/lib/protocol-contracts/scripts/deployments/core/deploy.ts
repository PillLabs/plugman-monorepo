import { network } from "hardhat";

import { isProtocolNetworkName } from "../../../lib/address.tools";
import { deployERC20Custody } from "./deploy-erc20-custody";
import { deployZetaConnector } from "./deploy-zeta-connector";
import { deployZetaToken } from "./deploy-zeta-token";

const networkName = network.name;

async function main() {
  if (!isProtocolNetworkName(networkName)) throw new Error("Invalid network name");

  const zetaTokenAddress = await deployZetaToken();
  await deployZetaConnector(zetaTokenAddress);
  await deployERC20Custody(zetaTokenAddress);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
