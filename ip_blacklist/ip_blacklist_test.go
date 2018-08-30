package ipblacklist

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hlts2/lilty"
)

func TestContains(t *testing.T) {
	tests := []struct {
		blacklistAddrs BlacklistAddrs
		addr           string
		expected       bool
	}{
		{
			blacklistAddrs: BlacklistAddrs{"192.168.33.10", "192.168.33.11"},
			addr:           "192.168.33.10",
			expected:       true,
		},
		{
			blacklistAddrs: BlacklistAddrs{"192.168.33.14"},
			addr:           "192.168.33.10",
			expected:       false,
		},
		{
			blacklistAddrs: BlacklistAddrs{},
			addr:           "192.168.33.10",
			expected:       false,
		},
	}

	for i, test := range tests {
		got := test.blacklistAddrs.Contains(test.addr)

		if test.expected != got {
			t.Errorf("tests[%d] - Contains is wrong. expected: %v, got: %v", i, test.expected, got)
		}
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		config   Config
		request  *http.Request
		data     []byte
		expected string
	}{
		{
			config: Config{
				BlacklistAddrs: []string{"192.168.33.10", "192.168.33.11"},
			},
			request: &http.Request{
				RemoteAddr: "192.168.33.10",
			},
			expected: "",
		},
		{
			config: Config{
				BlacklistAddrs: []string{"192.168.33.10", "192.168.33.11"},
			},
			request: &http.Request{
				RemoteAddr: "192.168.33.11",
			},
			expected: "",
		},
		{
			config: Config{
				BlacklistAddrs: []string{"192.168.33.10", "192.168.33.11"},
			},
			request: &http.Request{
				RemoteAddr: "192.168.33.15",
			},
			data:     []byte("192.168.33.15"),
			expected: "192.168.33.15",
		},
		{
			config: Config{
				BlacklistAddrs: []string{},
			},
			request: &http.Request{
				RemoteAddr: "192.168.33.15",
			},
			data:     []byte("192.168.33.15"),
			expected: "192.168.33.15",
		},
		{
			config: Config{
				BlacklistAddrs: []string{""},
			},
			request: &http.Request{
				RemoteAddr: "192.168.33.15",
			},
			data:     []byte("192.168.33.15"),
			expected: "192.168.33.15",
		},
	}

	for i, test := range tests {
		writer := httptest.NewRecorder()
		handler := New(test.config)(func(ctxt *lilty.Context) {
			ctxt.Write(200, test.data)
		})

		handler(&lilty.Context{
			Writer:  writer,
			Request: test.request,
		})

		got := writer.Body.String()

		if test.expected != got {
			t.Errorf("test[%d] - New middleware is wrong. expected: %v, got: %v", i, test.expected, got)
		}
	}
}
