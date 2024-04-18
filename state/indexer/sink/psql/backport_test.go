package psql

import (
	"github.com/KYVENetwork/cometbft/v37/state/indexer"
	"github.com/KYVENetwork/cometbft/v37/state/txindex"
)

var (
	_ indexer.BlockIndexer = BackportBlockIndexer{}
	_ txindex.TxIndexer    = BackportTxIndexer{}
)
