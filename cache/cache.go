package cache

import (
	"github.com/hlts2/gocache"
	"github.com/hlts2/lilty"
)

// New returns cache middle ware for lilty framework
func New() lilty.ChainHandler {
	_ = gocache.New()
	return func(next lilty.Handler) lilty.Handler {
		return func(ctxt *lilty.Context) {
			next(ctxt)
		}
	}
}
