package config

import (
	"time"
)

// Common is common config between full node and miner
type Common struct {
	API API
}

// API contains configs for API endpoint
type API struct {
	ListenAddress       string
	RemoteListenAddress string
	Timeout             Duration
}

// Duration is a wrapper type for time.Duration
// for decoding and encoding from/to TOML
type Duration time.Duration

// FullNode is a full node config
type FullNode struct {
	Common
	Metrics Metrics
}
type Metrics struct {
	Nickname   string
	HeadNotifs bool
}

func defCommon() Common {
	return Common{
		API: API{
			ListenAddress: "/ip4/127.0.0.1/tcp/3333/http",
			Timeout:       Duration(30 * time.Second),
		},
	}
}

// DefaultFullNode returns the default config
func DefaultFullNode() *FullNode {
	return &FullNode{
		Common: defCommon(),
	}
}
