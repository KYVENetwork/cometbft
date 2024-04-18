package test

import (
	"github.com/KYVENetwork/cometbft/v38/types"
)

// ConsensusParams returns a default set of ConsensusParams that are suitable
// for use in testing
func ConsensusParams() *types.ConsensusParams {
	c := types.DefaultConsensusParams()
	// enable vote extensions
	c.ABCI.VoteExtensionsEnableHeight = 1
	return c
}
