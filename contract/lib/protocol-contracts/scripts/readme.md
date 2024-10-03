## Worker script

To start localnet execute:

```
yarn compile
yarn localnet
```

This will run hardhat local node and worker script, which will deploy all contracts, and listen and react to events, facilitating communication between contracts.
Tasks to interact with localnet are located in `tasks/localnet`. To make use of default contract addresses on localnet, start localnet from scratch, so contracts are deployed on same addresses. Otherwise, provide custom addresses as tasks parameters.

To only show logs from worker and hide hardhat node logs:
```
yarn localnet --hide="NODE"
```