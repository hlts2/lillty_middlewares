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
		expected *http.Response
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
			expected: &http.Response{
				Status: "401 Unauthorized",
				Header: http.Header{
					"Www-Authenticate": {`Basic realm="Secret area"`},
				},
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
			expected: &http.Response{
				Status: "200 OK",
				Header: http.Header{},
			},
		},
	}

	for i, test := range tests {
		writer := httptest.NewRecorder()
		handler := New(test.config)(func(ctxt *lilty.Context) {})

		handler(&lilty.Context{
			Writer:  writer,
			Request: test.request,
		})

		response := writer.Result()

		if test.expected.Status != response.Status {
			t.Errorf("tests[%d] - New middleware Status is wrong. expected: %v, got: %v", i, test.expected.Status, response.Status)
		}

		if !reflect.DeepEqual(test.expected.Header, response.Header) {
			t.Errorf("tests[%d] - New middleware Header is wrong. expected: %v, got: %v", i, test.expected.Header, response.Header)
		}
	}

}
