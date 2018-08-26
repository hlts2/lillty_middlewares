package ipblacklist

import "testing"

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
			expected:       true,
		},
	}

	for i, test := range tests {
		got := test.blacklistAddrs.Contains(test.addr)

		if test.expected != got {
			t.Errorf("tests[%d] - Contains is wrong. expected: %v, got: %v", i, test.expected, got)
		}
	}
}
