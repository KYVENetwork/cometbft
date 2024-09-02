package mempool

import (
	memprotos "github.com/KYVENetwork/cometbft/v1/api/cometbft/mempool/v1"
	"github.com/KYVENetwork/cometbft/v1/types"
)

var (
	_ types.Wrapper   = &memprotos.Txs{}
	_ types.Unwrapper = &memprotos.Message{}
)
