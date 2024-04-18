package types

import cmtproto "github.com/KYVENetwork/cometbft/v034x/proto/cometbft/v034x/types"

// IsVoteTypeValid returns true if t is a valid vote type.
func IsVoteTypeValid(t cmtproto.SignedMsgType) bool {
	switch t {
	case cmtproto.PrevoteType, cmtproto.PrecommitType:
		return true
	default:
		return false
	}
}
