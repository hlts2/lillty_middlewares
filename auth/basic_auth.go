package auth

import (
	"github.com/hlts2/lilty"
)

// BasicAuth returns basic auth middleware for lilty framwwork
func BasicAuth(cUsername, cPassword string) lilty.ChainHandler {
	return func(handler lilty.Handler) lilty.Handler {
		return func(ctxt *lilty.Context) {
			username, password, ok := ctxt.Request.BasicAuth()

			match := cUsername == username && cPassword == password

			if !ok || !match {
				// TODO send error
				return
			}

			handler(ctxt)
		}
	}
}
