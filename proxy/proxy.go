package proxy

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/hlts2/lilty"
)

// Config represents configuration of proxy
type Config struct {
	Scheme string
	Host   string
}

// Proxy returns proxy middleware for lilty framework
func Proxy(c Config) lilty.ChainHandler {
	return func(next lilty.Handler) lilty.Handler {
		return func(ctxt *lilty.Context) {
			if ctxt.Scheme() != "https" || ctxt.Scheme() != "http" {
				log.Printf("not support scheme: %v\n", ctxt.Scheme())
				return
			}

			ctxt.Request.URL.Scheme = c.Scheme
			ctxt.Request.URL.Host = c.Host

			resp, err := http.DefaultTransport.RoundTrip(ctxt.Request)
			if err != nil {
				log.Println(err)
				return
			}

			defer resp.Body.Close()

			for _, cookie := range resp.Cookies() {
				http.SetCookie(ctxt.Writer, cookie)
			}

			copyHeader(ctxt.Writer, resp)

			b := readCloserToBytes(resp.Body)
			ctxt.Write(resp.StatusCode, b)
		}
	}
}

func copyHeader(writer http.ResponseWriter, resp *http.Response) {
	for key, values := range resp.Header {
		for _, value := range values {
			writer.Header().Add(key, value)
		}
	}
}

func readCloserToBytes(readCloser io.ReadCloser) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(readCloser)
	return buf.Bytes()
}
