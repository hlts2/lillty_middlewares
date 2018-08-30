package proxy

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
	"time"

	"github.com/hlts2/lilty"
)

func TestNew(t *testing.T) {
	expected := &http.Response{
		Status: "200 OK",
		Header: http.Header{
			"Content-Type":   []string{"application/json"},
			"Set-Cookie":     []string{"name=hiroto"},
			"Content-Length": []string{"0"},
			"Date":           []string{time.Unix(0, 0).String()},
		},
	}

	stop := runTestServer(":1234", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Date", time.Unix(0, 0).String())
		http.SetCookie(w, &http.Cookie{
			Name:  "name",
			Value: "hiroto",
		})
	}))

	writer := httptest.NewRecorder()

	config := Config{
		Scheme: "http",
		Host:   "localhost:1234",
	}

	handler := New(config)(func(ctxt *lilty.Context) {})

	handler(&lilty.Context{
		Writer: writer,
		Request: &http.Request{
			Method: http.MethodGet,
			URL: &url.URL{
				Scheme: "http",
				Path:   "/hoge",
			},
		},
	})

	response := writer.Result()

	if expected.Status != response.Status {
		t.Errorf("New Status is wrong. expected: %v, got: %v", expected.Status, response.Status)
	}

	if !reflect.DeepEqual(expected.Header, response.Header) {
		t.Errorf("New Header is wrong. expected: %v, got: %v", expected.Header, response.Header)
	}

	stop <- true
}

func runTestServer(addr string, handler http.Handler) chan bool {
	stop := make(chan bool, 1)

	s := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		<-stop
		s.Shutdown(context.Background())
	}()

	return stop
}
