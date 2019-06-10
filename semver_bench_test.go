package semver

import (
	"testing"
)

func BenchmarkVersionToString(b *testing.B) {
	version := Version{
		Major: 30,
		Minor: 5,
		Patch: 2019,
		Build: "766c5b54aae8d9de1708160817b93db1997e35cd",
	}
	for n := 0; n < b.N; n++ {
		version.String()
	}
}