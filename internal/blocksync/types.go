package blocksync

import (
	cmtbs "github.com/KYVENetwork/cometbft/v100/api/cometbft/blocksync/v1"
	"github.com/KYVENetwork/cometbft/v100/types"
)

var (
	_ types.Wrapper = &cmtbs.StatusRequest{}
	_ types.Wrapper = &cmtbs.StatusResponse{}
	_ types.Wrapper = &cmtbs.NoBlockResponse{}
	_ types.Wrapper = &cmtbs.BlockResponse{}
	_ types.Wrapper = &cmtbs.BlockRequest{}
)
