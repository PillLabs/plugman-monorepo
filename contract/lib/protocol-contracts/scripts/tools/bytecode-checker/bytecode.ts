import { BigNumber } from "ethers";
import { parseEther } from "ethers/lib/utils";

import {
  BNB_BSC,
  bscConnectorAddress,
  bscERC20CustodyAddress,
  bscTokenAddress,
  BTC_BTC,
  ETH_ETH,
  ethConnectorAddress,
  ethERC20CustodyAddress,
  ethTokenAddress,
} from "./bytecode.constants";
import {
  compareBytecode,
  encodeAddress,
  encodeNumber,
  getDeployedBytecode,
  getEtherscanBytecode,
  getZetaNodeBytecode,
  removeImmutableAddress,
  removeImmutableNumber,
} from "./bytecode.helpers";

const checkEthConnectorBytecode = async () => {
  const remoteBytecode = await getEtherscanBytecode("eth", ethConnectorAddress);
  const cleanRemoteBytecode = removeImmutableAddress(remoteBytecode, encodeAddress(ethTokenAddress));
  const deployedBytecode = await getDeployedBytecode("ZetaConnector.eth.sol/ZetaConnectorEth", "evm");
  compareBytecode(cleanRemoteBytecode, deployedBytecode);
};

const checkEthCustodyBytecode = async () => {
  const remoteBytecode = await getEtherscanBytecode("eth", ethERC20CustodyAddress);
  let cleanRemoteBytecode = removeImmutableAddress(remoteBytecode, encodeAddress(ethTokenAddress));
  cleanRemoteBytecode = removeImmutableNumber(cleanRemoteBytecode, encodeNumber(parseEther("1000"))); // zetaMaxFee
  const deployedBytecode = await getDeployedBytecode("ERC20Custody.sol/ERC20Custody", "evm");
  compareBytecode(cleanRemoteBytecode, deployedBytecode);
};

const checkBscConnectorBytecode = async () => {
  const remoteBytecode = await getEtherscanBytecode("bsc", bscConnectorAddress);
  const cleanRemoteBytecode = removeImmutableAddress(remoteBytecode, encodeAddress(bscTokenAddress));
  const deployedBytecode = await getDeployedBytecode("ZetaConnector.non-eth.sol/ZetaConnectorNonEth", "evm");
  compareBytecode(cleanRemoteBytecode, deployedBytecode);
};

const checkBscCustodyBytecode = async () => {
  const remoteBytecode = await getEtherscanBytecode("bsc", bscERC20CustodyAddress);
  let cleanRemoteBytecode = removeImmutableAddress(remoteBytecode, encodeAddress(bscTokenAddress));
  cleanRemoteBytecode = removeImmutableNumber(cleanRemoteBytecode, encodeNumber(parseEther("1000"))); // zetaMaxFee
  const deployedBytecode = await getDeployedBytecode("ERC20Custody.sol/ERC20Custody", "evm");
  compareBytecode(cleanRemoteBytecode, deployedBytecode);
};

const checkZRC20ETHBytecode = async () => {
  const remoteBytecode = await getZetaNodeBytecode(ETH_ETH);
  let cleanRemoteBytecode = removeImmutableNumber(remoteBytecode, encodeNumber(BigNumber.from("1"))); // ETH CHAIN ID
  cleanRemoteBytecode = removeImmutableNumber(cleanRemoteBytecode, encodeNumber(BigNumber.from("1"))); // Gas COIN TYPE
  const deployedBytecode = await getDeployedBytecode("ZRC20.sol/ZRC20", "zevm");
  compareBytecode(cleanRemoteBytecode, deployedBytecode);
};

const checkZRC20BTCBytecode = async () => {
  const remoteBytecode = await getZetaNodeBytecode(BTC_BTC);
  let cleanRemoteBytecode = removeImmutableNumber(remoteBytecode, encodeNumber(BigNumber.from("8332"))); // BTC CHAIN ID
  cleanRemoteBytecode = removeImmutableNumber(cleanRemoteBytecode, encodeNumber(BigNumber.from("1"))); // Gas COIN TYPE
  const deployedBytecode = await getDeployedBytecode("ZRC20.sol/ZRC20", "zevm");
  compareBytecode(cleanRemoteBytecode, deployedBytecode);
};

const checkZRC20BSCBytecode = async () => {
  const remoteBytecode = await getZetaNodeBytecode(BNB_BSC);
  let cleanRemoteBytecode = removeImmutableNumber(remoteBytecode, encodeNumber(BigNumber.from("56"))); // BSC CHAIN ID
  cleanRemoteBytecode = removeImmutableNumber(cleanRemoteBytecode, encodeNumber(BigNumber.from("1"))); // Gas COIN TYPE
  const deployedBytecode = await getDeployedBytecode("ZRC20.sol/ZRC20", "zevm");
  compareBytecode(cleanRemoteBytecode, deployedBytecode);
};

const checkBytecode = async () => {
  // ETH
  await checkEthConnectorBytecode();
  await checkEthCustodyBytecode();
  // BSC
  await checkBscConnectorBytecode();
  await checkBscCustodyBytecode();
  // ZEVM
  await checkZRC20ETHBytecode();
  await checkZRC20BTCBytecode();
  await checkZRC20BSCBytecode();
};

checkBytecode()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
