package cors

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/hlts2/lilty"
)

// Credentials is custom type for Access-Control-Allow-Credentials
type Credentials bool

// String returns "true" if c(Credentials) is true, otherwise it returns "false"
func (c Credentials) String() string {
	switch c {
	case true:
		return "true"
	case false:
		return "false"
	default:
		return "false"
	}
}

// Config represents configuration of cros
type Config struct {
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	AllowCredentials Credentials
	MaxAge           time.Duration
}

// New returns cors middleware for lilty framework
func New(c Config) lilty.ChainHandler {
	allowOrigins := strings.Join(c.AllowOrigins, ",")
	allowMethods := strings.Join(c.AllowMethods, ",")
	allowHeaders := strings.Join(c.AllowHeaders, ",")
	allowCredentials := c.AllowCredentials.String()
	maxAge := fmt.Sprint(c.MaxAge.Seconds())

	return func(next lilty.HandlerFunc) lilty.HandlerFunc {
		return func(ctxt *lilty.Context) {

			_, ok := ctxt.GetRequestHeader(lilty.AccessControlRequestMethod)

			// preflight
			if ok || ctxt.Request.Method == http.MethodOptions {
				ctxt.SetResponseHeader(lilty.AccessControlAllowOrigin, allowOrigins)
				ctxt.SetResponseHeader(lilty.AccessControlAllowMethods, allowMethods)
				ctxt.SetResponseHeader(lilty.AccessControlAllowHeaders, allowHeaders)
				ctxt.SetResponseHeader(lilty.AccessControlMaxAge, maxAge)
				return
			}

			ctxt.SetResponseHeader(lilty.AccessControlAllowOrigin, allowOrigins)
			ctxt.SetResponseHeader(lilty.AccessControlAllowCredentials, allowCredentials)

			next(ctxt)
		}
	}
}
