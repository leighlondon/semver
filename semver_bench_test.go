package semver

import (
	"testing"
)

func BenchmarkVersionToStringBuilder(b *testing.B) {
	version := Version{
		Major: 30,
		Minor: 5,
		Patch: 2019,
		Build: "766c5b54aae8d9de1708160817b93db1997e35cd",
	}
	for n := 0; n < b.N; n++ {
		version.stringBuilder()
	}
}

func BenchmarkVersionToStringBuilderFull(b *testing.B) {
	version := Version{
		Major: 30,
		Minor: 5,
		Patch: 2019,
		Build: "766c5b54aae8d9de1708160817b93db1997e35cd",
	}
	for n := 0; n < b.N; n++ {
		version.stringBuilderFull()
	}
}

func BenchmarkVersionToStringSprintf(b *testing.B) {
	version := Version{
		Major: 30,
		Minor: 5,
		Patch: 2019,
		Build: "766c5b54aae8d9de1708160817b93db1997e35cd",
	}
	for n := 0; n < b.N; n++ {
		version.sprintf()
	}
}
