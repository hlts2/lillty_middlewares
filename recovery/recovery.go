package recovery

import (
	"log"
	"net/http"

	"github.com/hlts2/lilty"
)

// New returns recovery middleware for lilty framework
func New() lilty.ChainHandler {
	return func(next lilty.HandlerFunc) lilty.HandlerFunc {
		return func(ctxt *lilty.Context) {
			defer func() {
				if err := recover(); err != nil {
					log.Printf("[PANIC RECOVER] error: %v", err)
					http.Error(ctxt.Writer, http.StatusText(500), 500)
				}
			}()

			next(ctxt)
		}
	}
}
