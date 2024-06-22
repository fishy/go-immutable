//go:build go1.23

package immutable_test

import (
	"fmt"
	"iter"
	"testing"

	"go.yhsif.com/immutable"
)

var _ iter.Seq2[int, string] = immutable.EmptyList[string]().All()

func BenchmarkListRangeOverFunc(b *testing.B) {
	for _, size := range sizes {
		b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
			orig := make([]int, size)
			for i := 0; i < size; i++ {
				orig[i] = i
			}
			b.Run("baseline", func(b *testing.B) {
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					for i, x := range orig {
						_ = i
						_ = x
					}
				}
			})
			b.Run("immutable", func(b *testing.B) {
				b.ReportAllocs()
				l := immutable.ListLiteral(orig...)
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					for i, x := range l.All() {
						_ = i
						_ = x
					}
				}
			})
		})
	}
}
