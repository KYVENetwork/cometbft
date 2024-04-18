package psql

import (
	"github.com/KYVENetwork/cometbft/v034x/state/indexer"
	"github.com/KYVENetwork/cometbft/v034x/state/txindex"
)

var (
	_ indexer.BlockIndexer = BackportBlockIndexer{}
	_ txindex.TxIndexer    = BackportTxIndexer{}
)
