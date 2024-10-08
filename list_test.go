package immutable_test

import (
	"fmt"
	"reflect"
	"testing"

	"go.yhsif.com/immutable"
)

func TestListBuilder(t *testing.T) {
	orig := []int{1, 2, 3}
	l := immutable.ListLiteral(orig...)
	size := l.Len()
	if size != 3 {
		t.Errorf("Len() expected 3, got %d", size)
	}
	list := make([]int, size)
	if err := l.Range(func(i int, x int) error {
		list[i] = x
		return nil
	}); err != nil {
		t.Errorf("Range should return nil, got: %v", err)
	}
	if !reflect.DeepEqual(orig, list) {
		t.Errorf("list expected %v, got %v", orig, list)
	}
}

func TestListBreak(t *testing.T) {
	index := 2
	orig := []int{1, 2, 3, 4, 5}
	short := orig[:index]
	l := immutable.ListLiteral(orig...)
	var list []int
	if err := l.Range(func(i int, x int) error {
		if i >= index {
			return immutable.ErrBreak
		}
		list = append(list, x)
		return nil
	}); err != immutable.ErrBreak {
		t.Errorf("Range should return: %v, got: %v", immutable.ErrBreak, err)
	}
	if !reflect.DeepEqual(short, list) {
		t.Errorf("list expected %v, got %v", short, list)
	}
}

func TestListString(t *testing.T) {
	orig := []int{1, 2, 3, 4, 5}
	l := immutable.ListLiteral(orig...)
	origStr := fmt.Sprintf("%v", orig)
	listStr := fmt.Sprintf("%v", l)
	if listStr != origStr {
		t.Errorf("List.String() expected %q, got %q", origStr, listStr)
	}
}

func TestListReslice(t *testing.T) {
	l := immutable.ListLiteral(0, 1, 2, 3, 4)
	reslice := l.Reslice(1, 3)
	if reslice.Len() != 2 {
		t.Errorf("Expected reslice length of 2, got %v", reslice)
	}
	if reslice.Get(0) != 1 || reslice.Get(1) != 2 {
		t.Errorf("Expected reslice of [1, 2], got %v", reslice)
	}
}

func TestEmptyList(t *testing.T) {
	if l := immutable.EmptyList[any]().Len(); l != 0 {
		t.Errorf("EmptyList.Len() expected 0, got %d", l)
	}
	if err := immutable.EmptyList[any]().Range(func(i int, x any) error {
		t.Errorf("EmptyList.Range called ListRangeFunc with %d, %v", i, x)
		return nil
	}); err != nil {
		t.Errorf("EmptyList.Range() returned error: %v", err)
	}
}

func BenchmarkListBuilder(b *testing.B) {
	b.Run("literal-10", func(b *testing.B) {
		b.Run("baseline", func(b *testing.B) {
			b.ReportAllocs()
			for range b.N {
				_ = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
			}
		})
		b.Run("immutable", func(b *testing.B) {
			b.ReportAllocs()
			for range b.N {
				immutable.ListLiteral(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
			}
		})
	},
	)
	for _, size := range sizes {
		b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
			b.Run("baseline", func(b *testing.B) {
				b.ReportAllocs()
				for range b.N {
					list := make([]int, size)
					for i := range size {
						list[i] = i
					}
				}
			})
			b.Run("immutable", func(b *testing.B) {
				b.ReportAllocs()
				for range b.N {
					list := make([]int, size)
					for i := 0; i < size; i++ {
						list[i] = i
					}
					immutable.ListLiteral(list...)
				}
			})
		})
	}
}

func BenchmarkListRange(b *testing.B) {
	for _, size := range sizes {
		b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
			orig := make([]int, size)
			for i := 0; i < size; i++ {
				orig[i] = i
			}
			b.Run("baseline", func(b *testing.B) {
				b.ReportAllocs()
				for range b.N {
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
				for range b.N {
					l.Range(func(i int, x int) error {
						return nil
					})
				}
			})
			b.Run("immutable-all", func(b *testing.B) {
				b.ReportAllocs()
				l := immutable.ListLiteral(orig...)
				b.ResetTimer()
				for range b.N {
					for i, x := range l.All() {
						_ = i
						_ = x
					}
				}
			})
		})
	}
}
