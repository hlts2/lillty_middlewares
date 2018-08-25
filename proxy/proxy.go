package proxy

import (
	"github.com/hlts2/lilty"
)

// Config represents configuration of proxy
type Config struct {
	Host string
}

// Proxy returns proxy middleware for lilty framework
func Proxy() lilty.ChainHandler {
	return func(next lilty.Handler) lilty.Handler {
		return func(ctxt *lilty.Context) {
			if ctxt.Scheme() != "https" || ctxt.Scheme() != "http" {
				// TODO error
				return
			}

			// TODO proxy
		}
	}
}
