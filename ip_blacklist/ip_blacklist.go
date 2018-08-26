package ipblacklist

import (
	"github.com/hlts2/lilty"
)

// BlacklistAddrs is custom type for blacklist address
type BlacklistAddrs []string

// Contains returns true if the target address is contained in the `RemoteAddrs`, false otherwise
func (b BlacklistAddrs) Contains(tgtAddr string) bool {
	if len(b) == 0 {
		return true
	}

	for _, addr := range b {
		if addr == tgtAddr {
			return true
		}
	}

	return false
}

// Config represents configuration of ip-blacklist middleware
type Config struct {
	BlacklistAddrs
}

// New returns ip-blacklist middlware of lilty framework
func New(c Config) lilty.ChainHandler {
	return func(next lilty.Handler) lilty.Handler {
		return func(ctxt *lilty.Context) {
			if c.Contains(ctxt.RemoteAddr()) {
				return
			}

			next(ctxt)
		}
	}
}
