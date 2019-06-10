package semver

import (
	"testing"
)

func BenchmarkVersionToString(b *testing.B) {
	v := Version{
		Major: 30,
		Minor: 5,
		Patch: 2019,
		Build: "766c5b54aae8d9de1708160817b93db1997e35cd",
	}

	b.Run("builder", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			v.builder()
		}
	})
	b.Run("builder2", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			v.builder2()
		}
	})
	b.Run("sprintf", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			v.sprintf()
		}
	})
	b.Run("sprintf2", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			v.sprintf2()
		}
	})
}
