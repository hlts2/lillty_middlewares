package redirect

import (
	"net/http"

	"github.com/hlts2/lilty"
)

// Code is custom type for status code
type Code int

// Check3xxx returns true if code is 3xx, otherwise it returns false
func (c Code) Check3xxx() bool {
	if c >= 300 && c < 400 {
		return true
	}
	return false
}

// Int converts Code type to int type
func (c Code) Int() int {
	return int(c)
}

// Config represents cofiguration of ssl redirect
type Config struct {
	Code
}

var defaultConfig = Config{
	Code: http.StatusMovedPermanently,
}

// Default returns ssl redirect middleware for lilty framework with default configuration
func Default() lilty.ChainHandler {
	return New(defaultConfig)
}

// New returns ssl redirect middleware for lilty framework
func New(c Config) lilty.ChainHandler {
	return func(next lilty.Handler) lilty.Handler {
		return func(ctxt *lilty.Context) {
			if !c.Code.Check3xxx() {
				// TODO log
				return
			}

			tgt := "https://" + ctxt.Host() + ctxt.Path()

			query := ctxt.Request.URL.RawQuery
			if len(query) > 0 {
				tgt += "?" + query
			}

			http.Redirect(ctxt.Writer, ctxt.Request, tgt, c.Code.Int())
		}
	}
}
