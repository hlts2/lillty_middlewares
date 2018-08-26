package cors

import (
	"github.com/hlts2/lilty"
)

// Config represents configuration of cros
type Config struct {
	headers map[string]string
}

// New returns cors middleware for lilty framework
func New() lilty.ChainHandler {
	return func(next lilty.Handler) lilty.Handler {
		return func(ctxt *lilty.Context) {
			next(ctxt)
		}
	}
}
