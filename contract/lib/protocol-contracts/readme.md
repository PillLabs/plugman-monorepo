### ⚠️ Important Notice: V2 in Development

We are currently developing Version 2 (V2) of our smart contract architecture. This new version will significantly enhance the developer experience for building Universal Apps.

Developers can already begin testing the new interface by referring to [the V2 Localnet guide](/v2_localnet.md).

# ZetaChain Protocol Contracts

This repository contains ZetaChain protocol contracts: Solidity source code,
generated Go bindings, deployed contract addresses and helper utilities.

## Importing Protocol Contracts

As a dApp developer, you can install the protocol contracts package into your
project:

```
yarn add --dev @zetachain/protocol-contracts
```

Getting the TSS address on BSC testnet:

```ts
import { getAddress } from "@zetachain/protocol-contracts";

getAddress("tss", "zeta_testnet");
```

Getting a ZRC-20 BSC USDT on ZetaChain Mainnet Beta:

```ts
import { getAddress } from "@zetachain/protocol-contracts";

getAddress("zrc20", "zeta_mainnet", "USDT.BSC");
```

The third argument (symbol) is only used when querying ZRC-20 addresses to
specify which token address is needed.

To view a table of all contracts visit the [Contract
Addresses](https://www.zetachain.com/docs/reference/contracts/) page in the
docs.

Importing
[`ZetaInterfaces`](https://www.zetachain.com/docs/developers/cross-chain-messaging/connector/)
and `ZetaInteractor` for cross-chain messaging:

```solidity
import "@zetachain/protocol-contracts/contracts/evm/interfaces/ZetaInterfaces.sol";
import "@zetachain/protocol-contracts/contracts/evm/tools/ZetaInteractor.sol";
```

Importing [ZRC20](https://www.zetachain.com/docs/developers/tokens/zrc20/) and
the [system
contract](https://www.zetachain.com/docs/developers/omnichain/system-contract/)
for omni-chain smart contracts:

```solidity
import "@zetachain/protocol-contracts/contracts/zevm/interfaces/IZRC20.sol";
import "@zetachain/protocol-contracts/contracts/zevm/interfaces/zContract.sol";
import "@zetachain/protocol-contracts/contracts/zevm/SystemContract.sol";
```

## Prerequisites for Development

Before you can contribute to this project, you must have the following
installed:

- [Node.js](https://nodejs.org/)
- [Yarn](https://yarnpkg.com/)
- [jq](https://stedolan.github.io/jq/)
- [abigen](https://geth.ethereum.org/docs/tools/abigen)
- [forge](https://getfoundry.sh/)

## Compiling Contracts

To compile the contracts, run the following command:

```
yarn compile
```

This will compile the Solidity contracts and output the resulting JSON artifacts
to the `artifacts` directory.

## Test

To run v1 tests (hardhat):

```
yarn test
```

To run v2 tests (hardhat):
```
yarn test:prototypes
```

To run v2 tests (forge):
```
forge test -vvvv
```

## Generating Go Bindings and Contract Addresses

To generate Go bindings for the Solidity contracts and fetch, run the following
command:

```
yarn generate
```

This will use `abigen` to generate Go bindings for the contracts and output the
resulting Go files to the `pkg` directory.

## Generating Contract Docs

To generate the docs from protocol contracts, run the following command:

```
yarn docs
```

This script uses `forge doc` to generate the docs.

## Contributing

If you would like to contribute to this project, please fork the repository and
submit a pull request. All contributions are welcome!
