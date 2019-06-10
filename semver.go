package semver // import "github.com/leighlondon/semver"

import (
	"fmt"
	"strconv"
	"strings"
)

// Version represents a semver version, with version numbers and optional
// build and prerelease metadata.
type Version struct {
	Major      int
	Minor      int
	Patch      int
	Build      string
	PreRelease string
}

func (v Version) String() string {
	return v.builder2()
}

func (v Version) builder() string {
	var ss strings.Builder
	ss.WriteString(strconv.Itoa(v.Major))
	ss.WriteRune('.')
	ss.WriteString(strconv.Itoa(v.Minor))
	ss.WriteRune('.')
	ss.WriteString(strconv.Itoa(v.Patch))
	if v.Build != "" {
		ss.WriteString("+" + v.Build)
	}
	if v.PreRelease != "" {
		ss.WriteString("-" + v.PreRelease)
	}
	return ss.String()
}

func (v Version) builder2() string {
	var ss strings.Builder
	ss.WriteString(strconv.Itoa(v.Major))
	ss.WriteRune('.')
	ss.WriteString(strconv.Itoa(v.Minor))
	ss.WriteRune('.')
	ss.WriteString(strconv.Itoa(v.Patch))
	if v.Build != "" {
		ss.WriteString("+")
		ss.WriteString(v.Build)
	}
	if v.PreRelease != "" {
		ss.WriteString("-")
		ss.WriteString(v.PreRelease)
	}
	return ss.String()
}

func (v Version) sprintf() string {
	s := fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
	if v.Build != "" {
		s = fmt.Sprintf("%s+%s", s, v.Build)
	}
	if v.PreRelease != "" {
		s = fmt.Sprintf("%s-%s", s, v.PreRelease)
	}
	return s
}

func (v Version) sprintf2() string {
	maj := strconv.Itoa(v.Major)
	min := strconv.Itoa(v.Minor)
	pat := strconv.Itoa(v.Patch)

	s := fmt.Sprintf("%s.%s.%s", maj, min, pat)
	if v.Build != "" {
		s = fmt.Sprintf("%s+%s", s, v.Build)
	}
	if v.PreRelease != "" {
		s = fmt.Sprintf("%s-%s", s, v.PreRelease)
	}
	return s
}
