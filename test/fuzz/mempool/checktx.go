package reactor

import (
	"github.com/KYVENetwork/cometbft/v100/abci/example/kvstore"
	"github.com/KYVENetwork/cometbft/v100/config"
	mempl "github.com/KYVENetwork/cometbft/v100/mempool"
	"github.com/KYVENetwork/cometbft/v100/proxy"
)

var mempool mempl.Mempool

func init() {
	app := kvstore.NewInMemoryApplication()
	cc := proxy.NewLocalClientCreator(app)
	appConnMem, _ := cc.NewABCIMempoolClient()
	err := appConnMem.Start()
	if err != nil {
		panic(err)
	}

	cfg := config.DefaultMempoolConfig()
	cfg.Broadcast = false
	mempool = mempl.NewCListMempool(cfg, appConnMem, 0)
}

func Fuzz(data []byte) int {
	_, err := mempool.CheckTx(data, "")
	if err != nil {
		return 0
	}

	return 1
}
