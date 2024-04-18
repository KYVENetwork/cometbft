package sr25519

import cmtjson "github.com/KYVENetwork/cometbft/v38/libs/json"

const (
	PrivKeyName = "cometbft/v38/PrivKeySr25519"
	PubKeyName  = "cometbft/v38/PubKeySr25519"
)

func init() {
	cmtjson.RegisterType(PubKey{}, PubKeyName)
	cmtjson.RegisterType(PrivKey{}, PrivKeyName)
}
