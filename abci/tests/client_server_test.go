package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	abciclient "github.com/KYVENetwork/cometbft/v38/abci/client"
	"github.com/KYVENetwork/cometbft/v38/abci/example/kvstore"
	abciserver "github.com/KYVENetwork/cometbft/v38/abci/server"
)

func TestClientServerNoAddrPrefix(t *testing.T) {
	t.Helper()

	addr := "localhost:26658"
	transport := "socket"
	app := kvstore.NewInMemoryApplication()

	server, err := abciserver.NewServer(addr, transport, app)
	assert.NoError(t, err, "expected no error on NewServer")
	err = server.Start()
	assert.NoError(t, err, "expected no error on server.Start")
	t.Cleanup(func() {
		if err := server.Stop(); err != nil {
			t.Error(err)
		}
	})

	client, err := abciclient.NewClient(addr, transport, true)
	assert.NoError(t, err, "expected no error on NewClient")
	err = client.Start()
	assert.NoError(t, err, "expected no error on client.Start")
	t.Cleanup(func() {
		if err := client.Stop(); err != nil {
			t.Error(err)
		}
	})
}
