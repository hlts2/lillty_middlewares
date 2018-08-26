package logger

import (
	"log"
	"time"

	"github.com/hlts2/lilty"
)

// New returns logging middleware for lilty framework.
// ie) x.x.x.x -- [2018-08-24 19:13:30 -700 JST] "GET / HTTP/1.1" "google.com" "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36" 70µs
func New() lilty.ChainHandler {
	return func(next lilty.Handler) lilty.Handler {
		return func(ctxt *lilty.Context) {
			start := time.Now()

			next(ctxt)

			log.Printf("%s -- [%s] \"%s %s %s\" \"%s\" \"%s\" %s\n",
				ctxt.RemoteAddr(),
				start.Format("2006-01-02 15:04:05 -700 MST"),
				ctxt.Method(), ctxt.Path(), ctxt.Proto(),
				ctxt.Host(),
				ctxt.UserAgent(),
				time.Since(start).String())
		}
	}
}
