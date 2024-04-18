package core

import (
	"errors"
	"fmt"

	ctypes "github.com/KYVENetwork/cometbft/v37/rpc/core/types"
	rpctypes "github.com/KYVENetwork/cometbft/v37/rpc/jsonrpc/types"
	"github.com/KYVENetwork/cometbft/v37/types"
)

// BroadcastEvidence broadcasts evidence of the misbehavior.
// More: https://docs.cometbft.com/v0.37/rpc/#/Evidence/broadcast_evidence
func BroadcastEvidence(ctx *rpctypes.Context, ev types.Evidence) (*ctypes.ResultBroadcastEvidence, error) {
	if ev == nil {
		return nil, errors.New("no evidence was provided")
	}

	if err := ev.ValidateBasic(); err != nil {
		return nil, fmt.Errorf("evidence.ValidateBasic failed: %w", err)
	}

	if err := env.EvidencePool.AddEvidence(ev); err != nil {
		return nil, fmt.Errorf("failed to add evidence: %w", err)
	}
	return &ctypes.ResultBroadcastEvidence{Hash: ev.Hash()}, nil
}
