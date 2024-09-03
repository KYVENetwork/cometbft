package statesync

import (
	ssproto "github.com/KYVENetwork/cometbft/v100/api/cometbft/statesync/v1"
	"github.com/KYVENetwork/cometbft/v100/types"
)

var (
	_ types.Wrapper = &ssproto.ChunkRequest{}
	_ types.Wrapper = &ssproto.ChunkResponse{}
	_ types.Wrapper = &ssproto.SnapshotsRequest{}
	_ types.Wrapper = &ssproto.SnapshotsResponse{}
)
