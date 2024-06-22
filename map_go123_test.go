//go:build go1.23

package immutable_test

import (
	"fmt"
	"iter"
	"testing"

	"go.yhsif.com/immutable"
)

var _ iter.Seq2[int, string] = immutable.EmptyMap[int, string]().All()

func BenchmarkMapRangeOverFunc(b *testing.B) {
	for _, size := range sizes {
		b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
			orig := make(map[int]int)
			for i := 0; i < size; i++ {
				orig[i] = i
			}
			b.Run("baseline", func(b *testing.B) {
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					for k, v := range orig {
						_ = k
						_ = v
					}
				}
			})
			b.Run("immutable", func(b *testing.B) {
				b.ReportAllocs()
				m := immutable.MapLiteral(orig)
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					for k, v := range m.All() {
						_ = k
						_ = v
					}
				}
			})
		})
	}
}
