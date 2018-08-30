package redirect

import "testing"

func TestCheck3xxx(t *testing.T) {
	tests := []struct {
		code     Code
		expected bool
	}{
		{
			code:     314,
			expected: true,
		},
		{
			code:     300,
			expected: true,
		},
		{
			code:     299,
			expected: false,
		},
		{
			code:     400,
			expected: false,
		},
	}

	for i, test := range tests {
		got := test.code.Check3xx()

		if test.expected != got {
			t.Errorf("tests[%d] - Check3xxx is wrong. expected: %v, got: %v", i, test.expected, got)
		}
	}
}
