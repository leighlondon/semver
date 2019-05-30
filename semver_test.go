package semver

import (
	"testing"
)

func TestStringFormatting(t *testing.T) {
	samples := []struct {
		version  Version
		expected string
	}{
		{Version{}, "0.0.0"},
		{Version{Major: 0, Minor: 0, Patch: 0}, "0.0.0"},
		{Version{1, 0, 4, "abcd", "hmm"}, "1.0.4+abcd-hmm"},
		{Version{1, 0x8, 4, "abcd", "hmm"}, "1.8.4+abcd-hmm"},
	}

	for _, s := range samples {
		if actual := s.version.String(); actual != s.expected {
			t.Errorf("got='%s' expected='%s' %#v",
				actual, s.expected, s.version)
		}
	}
}
