## Setup echidna

```
brew install echidna
solc-select use 0.8.7
```

## Execute contract tests

```
echidna test/fuzz/ERC20CustodyNewEchidnaTest.sol --contract ERC20CustodyNewEchidnaTest  --config echidna.yaml
echidna test/fuzz/GatewayEVMEchidnaTest.sol --contract GatewayEVMEchidnaTest  --config echidna.yaml
```