import { parseEther } from "ethers/lib/utils";

export const ZETA_INITIAL_SUPPLY = 2_100_000_000;

export const MAX_ETH_ADDRESS = "0xffffffffffffffffffffffffffffffffffffffff";

// dev: this values should be calculated using get-salt script
const SALT_NUMBERS = {
  polygon_mainnet: {
    zetaConnector: "0",
    zetaConsumer: "0",
    zetaERC20Custody: "0",
    zetaToken: "0",
  },
  baobab_testnet: {
    zetaConnector: "71733",
    zetaConsumer: "0",
    zetaERC20Custody: "195084",
    zetaToken: "29265",
  },
  bsc_mainnet: {
    zetaConnector: "71733",
    zetaConsumer: "0",
    zetaERC20Custody: "195084",
    zetaToken: "29265",
  },
  bsc_testnet: {
    zetaConnector: "71733",
    zetaConsumer: "0",
    zetaERC20Custody: "195084",
    zetaToken: "29265",
  },
  btc_testnet: {
    zetaConnector: "",
    zetaConsumer: "",
    zetaERC20Custody: "",
    zetaToken: "",
  },
  eth_mainnet: {
    zetaConnector: "84286",
    zetaConsumer: "0",
    zetaERC20Custody: "926526",
    zetaToken: "0",
  },
  goerli_testnet: {
    zetaConnector: "1414",
    zetaConsumer: "0",
    zetaERC20Custody: "87967",
    zetaToken: "84108",
  },
  mumbai_testnet: {
    zetaConnector: "71733",
    zetaConsumer: "0",
    zetaERC20Custody: "195084",
    zetaToken: "29265",
  },
  sepolia_testnet: {
    zetaConnector: "1414",
    zetaConsumer: "0",
    zetaERC20Custody: "87967",
    zetaToken: "84108",
  },
  zeta_testnet: {
    zetaConnector: "71733",
    zetaConsumer: "0",
    zetaERC20Custody: "195084",
    zetaToken: "29265",
  },
};

export const getSaltNumber = (contractName: string, networkName: string) => {
  const saltNumber = SALT_NUMBERS[networkName][contractName];

  if (!saltNumber) {
    throw new Error(`Salt number for ${contractName} on ${networkName} is not defined.`);
  }

  return saltNumber;
};

export const ERC20_CUSTODY_ZETA_FEE = "0";
export const ERC20_CUSTODY_ZETA_MAX_FEE = parseEther("1000");
