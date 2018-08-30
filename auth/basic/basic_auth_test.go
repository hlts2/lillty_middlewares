package auth

import (
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/hlts2/lilty"
)

func TestNew(t *testing.T) {
	tests := []struct {
		config   Config
		request  *http.Request
		expected http.Header
	}{
		{
			config: Config{
				Username: "name",
				Password: "pass",
				Realm:    "Secret area",
			},
			request: &http.Request{
				Header: http.Header{
					"Authorization": {""},
				},
			},
			expected: http.Header{
				"Www-Authenticate": {`Basic realm="Secret area"`},
			},
		},
		{
			config: Config{
				Username: "name",
				Password: "pass",
				Realm:    "Secret area",
			},
			request: &http.Request{
				Header: http.Header{
					"Authorization": []string{"Basic " + base64.StdEncoding.EncodeToString([]byte("name:pass"))},
				},
			},
			expected: http.Header{},
		},
	}

	for i, test := range tests {
		writer := httptest.NewRecorder()
		handler := New(test.config)(func(ctxt *lilty.Context) {})

		handler(&lilty.Context{
			Writer:  writer,
			Request: test.request,
		})

		got := writer.Header()
		if !reflect.DeepEqual(test.expected, got) {
			t.Errorf("tests[%d] - New middleware is wrong. expected: %v, got: %v", i, test.expected, got)
		}
	}
}
