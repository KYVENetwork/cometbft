package reactor

import (
	"github.com/KYVENetwork/cometbft/v38/abci/example/kvstore"
	"github.com/KYVENetwork/cometbft/v38/config"
	mempl "github.com/KYVENetwork/cometbft/v38/mempool"
	"github.com/KYVENetwork/cometbft/v38/proxy"
)

var mempool mempl.Mempool

func init() {
	app := kvstore.NewInMemoryApplication()
	cc := proxy.NewLocalClientCreator(app)
	appConnMem, _ := cc.NewABCIClient()
	err := appConnMem.Start()
	if err != nil {
		panic(err)
	}

	cfg := config.DefaultMempoolConfig()
	cfg.Broadcast = false
	mempool = mempl.NewCListMempool(cfg, appConnMem, 0)
}

func Fuzz(data []byte) int {
	err := mempool.CheckTx(data, nil, mempl.TxInfo{})
	if err != nil {
		return 0
	}

	return 1
}
