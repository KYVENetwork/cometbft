package psql

import (
	"github.com/KYVENetwork/cometbft/v100/state/indexer"
	"github.com/KYVENetwork/cometbft/v100/state/txindex"
)

var (
	_ indexer.BlockIndexer = BackportBlockIndexer{}
	_ txindex.TxIndexer    = BackportTxIndexer{}
)
