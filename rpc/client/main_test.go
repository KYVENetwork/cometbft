package client_test

import (
	"os"
	"testing"

	"github.com/KYVENetwork/cometbft/v100/abci/example/kvstore"
	nm "github.com/KYVENetwork/cometbft/v100/node"
	rpctest "github.com/KYVENetwork/cometbft/v100/rpc/test"
)

var node *nm.Node

func TestMain(m *testing.M) {
	// start a CometBFT node (and kvstore) in the background to test-2 against
	dir, err := os.MkdirTemp("/tmp", "rpc-client-test-2")
	if err != nil {
		panic(err)
	}

	app := kvstore.NewPersistentApplication(dir)
	// If testing block event generation
	// app.SetGenBlockEvents() // needs to be called here (see TestBlockSearch in rpc_test.go)
	node = rpctest.StartCometBFT(app)

	code := m.Run()

	// and shut down proper at the end
	rpctest.StopCometBFT(node)
	_ = os.RemoveAll(dir)
	os.Exit(code)
}
