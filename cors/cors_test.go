package cors

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/hlts2/lilty"
)

func TestStrng(t *testing.T) {
	tests := []struct {
		credentials Credentials
		expexted    string
	}{
		{
			credentials: true,
			expexted:    "true",
		},
		{
			credentials: false,
			expexted:    "false",
		},
	}

	for i, test := range tests {
		got := test.credentials.String()
		if test.expexted != got {
			t.Errorf("tests[%d] - String is wrong. expexted: %v, got: %v", i, test.expexted, got)
		}
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		config   Config
		request  *http.Request
		expected http.Header
	}{
		// preflight
		{
			config: Config{
				AllowOrigins:     []string{"192.168.33.10"},
				AllowMethods:     []string{"GET", "POST"},
				AllowHeaders:     []string{"Content-Type"},
				AllowCredentials: true,
				MaxAge:           10 * time.Second,
			},
			request: &http.Request{
				Method: http.MethodOptions,
			},
			expected: http.Header{
				"Access-Control-Allow-Origin":  {"192.168.33.10"},
				"Access-Control-Allow-Methods": {"GET,POST"},
				"Access-Control-Allow-Headers": {"Content-Type"},
				"Access-Control-Max-Age":       {"10"},
			},
		},
		{
			config: Config{
				AllowOrigins:     []string{"192.168.33.10"},
				AllowMethods:     []string{"GET", "POST"},
				AllowHeaders:     []string{"Content-Type"},
				AllowCredentials: true,
				MaxAge:           10 * time.Second,
			},
			request: &http.Request{
				Header: http.Header{
					"Access-Control-Request-Method": {"GET"},
				},
			},
			expected: http.Header{
				"Access-Control-Allow-Origin":  {"192.168.33.10"},
				"Access-Control-Allow-Methods": {"GET,POST"},
				"Access-Control-Allow-Headers": {"Content-Type"},
				"Access-Control-Max-Age":       {"10"},
			},
		},

		// not preflight
		{
			config: Config{
				AllowOrigins:     []string{"192.168.33.10"},
				AllowCredentials: true,
			},
			request: &http.Request{
				Method: "GET",
			},
			expected: http.Header{
				"Access-Control-Allow-Origin":      {"192.168.33.10"},
				"Access-Control-Allow-Credentials": {"true"},
			},
		},
	}

	for _, test := range tests {
		writer := httptest.NewRecorder()
		handler := New(test.config)(func(ctxt *lilty.Context) {})

		handler(&lilty.Context{
			Writer:  writer,
			Request: test.request,
		})

		if !reflect.DeepEqual(test.expected, writer.Header()) {
			t.Errorf("New middleware is wrong. expected: %v, got: %v", test.expected, writer.Header())
		}
	}
}
