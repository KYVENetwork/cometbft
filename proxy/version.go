package proxy

import (
	abci "github.com/KYVENetwork/cometbft/v37/abci/types"
	"github.com/KYVENetwork/cometbft/v37/version"
)

// RequestInfo contains all the information for sending
// the abci.RequestInfo message during handshake with the app.
// It contains only compile-time version information.
var RequestInfo = abci.RequestInfo{
	Version:      version.TMCoreSemVer,
	BlockVersion: version.BlockProtocol,
	P2PVersion:   version.P2PProtocol,
	AbciVersion:  version.ABCIVersion,
}
