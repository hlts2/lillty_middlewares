package auth

import (
	"fmt"

	"github.com/hlts2/lilty"
)

// Config represents configuration of basic auth
type Config struct {
	Username string
	Password string
	Realm    string
}

// BasicAuth returns basic auth middleware for lilty framwwork
func BasicAuth(c Config) lilty.ChainHandler {
	return func(handler lilty.Handler) lilty.Handler {
		return func(ctxt *lilty.Context) {
			username, password, ok := ctxt.Request.BasicAuth()

			match := c.Username == username && c.Password == password

			if !ok || !match {
				ctxt.SetResponseHeader(lilty.WWWAuthenticate, fmt.Sprintf(`Basic realm="%s"`, c.Realm))
				ctxt.SetStatusCode(401)
				return
			}

			handler(ctxt)
		}
	}
}
