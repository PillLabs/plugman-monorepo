import { network } from "hardhat";

import { isProtocolNetworkName } from "../../../lib/address.tools";
import { deterministicDeployERC20Custody } from "./deterministic-deploy-erc20-custody";
import { deterministicDeployZetaConnector } from "./deterministic-deploy-zeta-connector";
import { deterministicDeployZetaToken } from "./deterministic-deploy-zeta-token";

const networkName = network.name;

async function main() {
  if (!isProtocolNetworkName(networkName)) throw new Error("Invalid network name");

  await deterministicDeployZetaToken();
  await deterministicDeployZetaConnector();
  await deterministicDeployERC20Custody();
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
