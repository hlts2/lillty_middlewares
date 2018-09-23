package ipblacklist

import (
	"github.com/hlts2/lilty"
)

// BlacklistAddrs is custom type for blacklist addresses
type BlacklistAddrs []string

// Contains returns true if the remote address is contained in the `BlacklistAddrs`, false otherwise
func (b BlacklistAddrs) Contains(tgtAddr string) bool {
	if len(b) == 0 {
		return false
	}

	for _, addr := range b {
		if addr == tgtAddr {
			return true
		}
	}

	return false
}

// New returns ip-blacklist middlware of lilty framework
func New(addr BlacklistAddrs) lilty.ChainHandler {
	return func(next lilty.HandlerFunc) lilty.HandlerFunc {
		return func(ctxt *lilty.Context) {
			if addr.Contains(ctxt.Request.RemoteAddr) {
				return
			}

			next(ctxt)
		}
	}
}
