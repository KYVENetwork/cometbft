package block

import (
	"errors"
	"fmt"

	"github.com/KYVENetwork/cometbft/v100/config"
	"github.com/KYVENetwork/cometbft/v100/state/indexer"
	blockidxkv "github.com/KYVENetwork/cometbft/v100/state/indexer/block/kv"
	blockidxnull "github.com/KYVENetwork/cometbft/v100/state/indexer/block/null"
	"github.com/KYVENetwork/cometbft/v100/state/indexer/sink/psql"
	"github.com/KYVENetwork/cometbft/v100/state/txindex"
	"github.com/KYVENetwork/cometbft/v100/state/txindex/kv"
	"github.com/KYVENetwork/cometbft/v100/state/txindex/null"
	dbm "github.com/cometbft/cometbft-db"
)

// EventSinksFromConfig constructs a slice of indexer.EventSink using the provided
// configuration.
func IndexerFromConfig(cfg *config.Config, dbProvider config.DBProvider, chainID string) (
	txIdx txindex.TxIndexer, blockIdx indexer.BlockIndexer, allIndexersDisabled bool, err error,
) {
	switch cfg.TxIndex.Indexer {
	case "kv":
		store, err := dbProvider(&config.DBContext{ID: "tx_index", Config: cfg})
		if err != nil {
			return nil, nil, false, err
		}

		return kv.NewTxIndex(store),
			blockidxkv.New(dbm.NewPrefixDB(store, []byte("block_events")),
				blockidxkv.WithCompaction(cfg.Storage.Compact, cfg.Storage.CompactionInterval)),
			false,
			nil

	case "psql":
		conn := cfg.TxIndex.PsqlConn
		if conn == "" {
			return nil, nil, false, errors.New("the psql connection settings cannot be empty")
		}
		es, err := psql.NewEventSink(cfg.TxIndex.PsqlConn, chainID)
		if err != nil {
			return nil, nil, false, fmt.Errorf("creating psql indexer: %w", err)
		}
		return es.TxIndexer(), es.BlockIndexer(), false, nil

	default:
		return &null.TxIndex{}, &blockidxnull.BlockerIndexer{}, true, nil
	}
}
