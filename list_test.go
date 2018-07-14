package immutable_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/fishy/go-immutable"
)

func TestListBuilder(t *testing.T) {
	orig := []interface{}{1, 2, 3}
	l := immutable.ListLiteral(orig...)
	size := l.Len()
	if size != 3 {
		t.Errorf("Len() expected 3, got %d", size)
	}
	list := make([]interface{}, size)
	if err := l.Range(func(i int, x interface{}) error {
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
	orig := []interface{}{1, 2, 3, 4, 5}
	short := orig[:index]
	l := immutable.ListLiteral(orig...)
	var list []interface{}
	if err := l.Range(func(i int, x interface{}) error {
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
	orig := []interface{}{1, 2, 3, 4, 5}
	l := immutable.ListLiteral(orig...)
	origStr := fmt.Sprintf("%v", orig)
	listStr := fmt.Sprintf("%v", l)
	if listStr != origStr {
		t.Errorf("List.String() expected %q, got %q", origStr, listStr)
	}
}

func BenchmarkListBuilder(b *testing.B) {
	b.Run(
		"literal-10",
		func(b *testing.B) {
			b.Run(
				"baseline",
				func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						_ = []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
					}
				},
			)
			b.Run(
				"immutable",
				func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						immutable.ListLiteral(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
					}
				},
			)
		},
	)
	for _, size := range sizes {
		b.Run(
			fmt.Sprintf("%d", size),
			func(b *testing.B) {
				b.Run(
					"baseline",
					func(b *testing.B) {
						for i := 0; i < b.N; i++ {
							list := make([]interface{}, size)
							for j := 0; j < size; j++ {
								list[j] = j
							}
						}
					},
				)
				b.Run(
					"immutable",
					func(b *testing.B) {
						for i := 0; i < b.N; i++ {
							list := make([]interface{}, size)
							for j := 0; j < size; j++ {
								list[j] = j
							}
							immutable.ListLiteral(list...)
						}
					},
				)
			},
		)
	}
}

func BenchmarkListRange(b *testing.B) {
	for _, size := range sizes {
		b.Run(
			fmt.Sprintf("%d", size),
			func(b *testing.B) {
				orig := make([]interface{}, size)
				for i := 0; i < size; i++ {
					orig[i] = i
				}
				b.Run(
					"baseline",
					func(b *testing.B) {
						for i := 0; i < b.N; i++ {
							for range orig {
							}
						}
					},
				)
				b.Run(
					"immutable",
					func(b *testing.B) {
						l := immutable.ListLiteral(orig...)
						b.ResetTimer()
						for i := 0; i < b.N; i++ {
							l.Range(func(i int, x interface{}) error {
								return nil
							})
						}
					},
				)
			},
		)
	}
}
