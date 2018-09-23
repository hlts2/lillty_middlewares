package recovery

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hlts2/lilty"
)

func TestNew(t *testing.T) {
	tests := []struct {
		handler  lilty.HandlerFunc
		expected *http.Response
	}{
		{
			handler: New()(func(ctxt *lilty.Context) {
				panic("panic occured")
			}),
			expected: &http.Response{
				StatusCode: 500,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("Internal Server Error\n"))),
			},
		},
		{
			handler: New()(func(ctxt *lilty.Context) {
				ctxt.Write(200, []byte("hello world"))
			}),
			expected: &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("hello world"))),
			},
		},
	}

	for i, test := range tests {
		writer := httptest.NewRecorder()

		test.handler(&lilty.Context{
			Writer: writer,
		})

		if test.expected.StatusCode != writer.Code {
			t.Errorf("tests[%d] - New Code is wrong. expected: %v, got: %v", i, test.expected.Status, writer.Code)
		}

		buf := new(bytes.Buffer)
		buf.ReadFrom(test.expected.Body)

		if buf.String() != writer.Body.String() {
			t.Errorf("tests[%d] - New Body is wrong. expected: %v, got: %v", i, buf.String(), writer.Body.String())
		}
	}
}
