package evidence

import (
	"github.com/KYVENetwork/cometbft/v034x/types"
)

//go:generate ../scripts/mockery_generate.sh BlockStore

type BlockStore interface {
	LoadBlockMeta(height int64) *types.BlockMeta
	LoadBlockCommit(height int64) *types.Commit
	Height() int64
}
