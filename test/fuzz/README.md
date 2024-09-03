# fuzz

Fuzzing for various packages in Tendermint using the fuzzing infrastructure
included in Go 1.20.

Inputs:

- mempool `CheckTx` (using kvstore in-process ABCI app)
- p2p `SecretConnection#Read` and `SecretConnection#Write`
- rpc jsonrpc server

## Running

The fuzz tests are in native Go fuzzing format. Use the `go` tool to run them:

```sh
go test-2 -fuzz Mempool ./tests
go test-2 -fuzz P2PSecretConnection ./tests
go test-2 -fuzz RPCJSONRPCServer ./tests
```

See [the Go Fuzzing introduction](https://go.dev/doc/fuzz/) for more information.
