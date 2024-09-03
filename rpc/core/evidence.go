package core

import (
	"errors"
	"fmt"

	ctypes "github.com/KYVENetwork/cometbft/v100/rpc/core/types"
	rpctypes "github.com/KYVENetwork/cometbft/v100/rpc/jsonrpc/types"
	"github.com/KYVENetwork/cometbft/v100/types"
)

// BroadcastEvidence broadcasts evidence of the misbehavior.
// More: https://docs.cometbft.com/main/rpc/#/Evidence/broadcast_evidence
func (env *Environment) BroadcastEvidence(
	_ *rpctypes.Context,
	ev types.Evidence,
) (*ctypes.ResultBroadcastEvidence, error) {
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
