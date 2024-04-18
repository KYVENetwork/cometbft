package sr25519

import (
	"github.com/KYVENetwork/cometbft/v34/crypto"
	cmtjson "github.com/KYVENetwork/cometbft/v34/libs/json"
)

var _ crypto.PrivKey = PrivKey{}

const (
	PrivKeyName = "cometbft/PrivKeySr25519"
	PubKeyName  = "cometbft/PubKeySr25519"

	// SignatureSize is the size of an Edwards25519 signature. Namely the size of a compressed
	// Sr25519 point, and a field element. Both of which are 32 bytes.
	SignatureSize = 64
)

func init() {

	cmtjson.RegisterType(PubKey{}, PubKeyName)
	cmtjson.RegisterType(PrivKey{}, PrivKeyName)
}
