package semver

import (
	"testing"
)

func BenchmarkVersionToString(b *testing.B) {
	cases := []struct {
		name string
		v    Version
	}{
		{"full", Version{Major: 30, Minor: 5, Patch: 2019, Build: "766c5b54aae8d9de1708160817b93db1997e35cd"}},
		{"classic", Version{Major: 1, Minor: 2, Patch: 3}},
	}

	for _, c := range cases {
		b.Run("builder/"+c.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				c.v.builder()
			}
		})
		b.Run("builder2/"+c.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				c.v.builder2()
			}
		})
		b.Run("sprintf/"+c.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				c.v.sprintf()
			}
		})
		b.Run("sprintf2/"+c.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				c.v.sprintf2()
			}
		})
	}
}
