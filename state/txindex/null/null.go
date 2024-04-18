package null

import (
	"context"
	"errors"

	"github.com/KYVENetwork/cometbft/v38/libs/log"

	abci "github.com/KYVENetwork/cometbft/v38/abci/types"
	"github.com/KYVENetwork/cometbft/v38/libs/pubsub/query"
	"github.com/KYVENetwork/cometbft/v38/state/txindex"
)

var _ txindex.TxIndexer = (*TxIndex)(nil)

// TxIndex acts as a /dev/null.
type TxIndex struct{}

// Get on a TxIndex is disabled and panics when invoked.
func (txi *TxIndex) Get(_ []byte) (*abci.TxResult, error) {
	return nil, errors.New(`indexing is disabled (set 'tx_index = "kv"' in config)`)
}

// AddBatch is a noop and always returns nil.
func (txi *TxIndex) AddBatch(_ *txindex.Batch) error {
	return nil
}

// Index is a noop and always returns nil.
func (txi *TxIndex) Index(_ *abci.TxResult) error {
	return nil
}

func (txi *TxIndex) Search(_ context.Context, _ *query.Query) ([]*abci.TxResult, error) {
	return []*abci.TxResult{}, nil
}

func (txi *TxIndex) SetLogger(log.Logger) {

}
