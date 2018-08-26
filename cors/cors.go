package cors

import (
	"strings"
	"time"

	"github.com/hlts2/lilty"
)

// Credentials is custom type for Access-Control-Allow-Credentials
type Credentials bool

// String returns "true" if c is true, otherwise it returns "false"
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
	_ = strings.Join(c.AllowOrigins, ",")
	_ = strings.Join(c.AllowMethods, ",")
	_ = strings.Join(c.AllowHeaders, ",")
	_ = c.AllowCredentials.String()

	return func(next lilty.Handler) lilty.Handler {
		return func(ctxt *lilty.Context) {
			// TODO 処理を追加
		}
	}
}
