//go:build gofuzz || go1.20

package tests

import (
	"testing"

	abciclient "github.com/KYVENetwork/cometbft/v100/abci/client"
	"github.com/KYVENetwork/cometbft/v100/abci/example/kvstore"
	"github.com/KYVENetwork/cometbft/v100/config"
	cmtsync "github.com/KYVENetwork/cometbft/v100/libs/sync"
	mempl "github.com/KYVENetwork/cometbft/v100/mempool"
)

func FuzzMempool(f *testing.F) {
	app := kvstore.NewInMemoryApplication()
	mtx := new(cmtsync.Mutex)
	conn := abciclient.NewLocalClient(mtx, app)
	err := conn.Start()
	if err != nil {
		panic(err)
	}

	cfg := config.DefaultMempoolConfig()
	cfg.Broadcast = false

	mp := mempl.NewCListMempool(cfg, conn, 0)

	f.Fuzz(func(_ *testing.T, data []byte) {
		_, _ = mp.CheckTx(data, "")
	})
}
