package auth

import (
	"github.com/hlts2/lilty"
)

// Config represents configuration of basic auth
type Config struct {
	Username string
	Password string
	Realm    string
}

// New returns basic auth middleware for lilty framewwork
func New(c Config) lilty.ChainHandler {
	return func(next lilty.HandlerFunc) lilty.HandlerFunc {
		return func(ctxt *lilty.Context) {
			username, password, ok := ctxt.Request.BasicAuth()

			match := c.Username == username && c.Password == password

			if !ok || !match {
				ctxt.SetResponseHeader(lilty.WWWAuthenticate, `Basic realm="`+c.Realm+`"`)
				ctxt.SetStatusCode(401)
				return
			}

			next(ctxt)
		}
	}
}
