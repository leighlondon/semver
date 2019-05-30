package semver // import "github.com/leighlondon/semver"

import (
	"strconv"
	"strings"
)

// Version ...
type Version struct {
	Major      int
	Minor      int
	Patch      int
	Build      string
	PreRelease string
}

func (v Version) String() string {
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
