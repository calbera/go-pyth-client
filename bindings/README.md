# Bindings

This directory contains the bindings for the Pyth EVM contracts. To generate the bindings, run `make bindings` from the root directory. This requires the Pyth EVM contracts' ABIs to be built and available in the `bindings/abis` sub-directory.

## Building

To get the necessary output files, follow the instructions in the [Pyth Solidity SDK](https://github.com/pyth-network/pyth-crosschain/tree/main/target_chains/ethereum/sdk/solidity).

To generate using abigen, follow the instructions in [Go-Ethereum's abigen](https://geth.ethereum.org/docs/tools/abigen#main-content) and add the `go:generate` command in `gen.go`.
