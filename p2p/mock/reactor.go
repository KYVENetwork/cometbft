package mock

import (
	"github.com/KYVENetwork/cometbft/v1/libs/log"
	"github.com/KYVENetwork/cometbft/v1/p2p"
	"github.com/KYVENetwork/cometbft/v1/p2p/conn"
)

type Reactor struct {
	p2p.BaseReactor

	Channels []*conn.ChannelDescriptor
}

func NewReactor() *Reactor {
	r := &Reactor{}
	r.BaseReactor = *p2p.NewBaseReactor("Mock-PEX", r)
	r.SetLogger(log.TestingLogger())
	return r
}

func (r *Reactor) GetChannels() []*conn.ChannelDescriptor { return r.Channels }
func (*Reactor) AddPeer(_ p2p.Peer)                       {}
func (*Reactor) RemovePeer(_ p2p.Peer, _ any)             {}
func (*Reactor) Receive(_ p2p.Envelope)                   {}
