package semver

import (
	"fmt"
	"testing"
)

func TestStringConversion(t *testing.T) {
	v := Version{Major: 1}
	ex := "1.0.0"
	if s := fmt.Sprintf("%s", v); s != ex {
		t.Errorf("expected formatting, got '%s'", s)
	}
}

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
		t.Run("builder/"+s.expected, func(t *testing.T) {
			if actual := s.version.builder(); actual != s.expected {
				t.Errorf("got='%s' expected='%s' %#v",
					actual, s.expected, s.version)
			}
		})
		t.Run("builder2/"+s.expected, func(t *testing.T) {
			if actual := s.version.builder2(); actual != s.expected {
				t.Errorf("got='%s' expected='%s' %#v",
					actual, s.expected, s.version)
			}
		})
		t.Run("sprintf/"+s.expected, func(t *testing.T) {
			if actual := s.version.sprintf(); actual != s.expected {
				t.Errorf("got='%s' expected='%s' %#v",
					actual, s.expected, s.version)
			}
		})
		t.Run("sprintf2/"+s.expected, func(t *testing.T) {
			if actual := s.version.sprintf2(); actual != s.expected {
				t.Errorf("got='%s' expected='%s' %#v",
					actual, s.expected, s.version)
			}
		})
	}
}
