package recovery

import (
	"log"

	"github.com/hlts2/lilty"
)

// New returns recovery middleware for lilty framework
func New() lilty.ChainHandler {
	return func(next lilty.Handler) lilty.Handler {
		return func(ctxt *lilty.Context) {
			defer func() {
				if err := recover(); err != nil {
					log.Printf("[PANIC RECOVER] error: %v", err)
				}
			}()

			next(ctxt)
		}
	}
}
