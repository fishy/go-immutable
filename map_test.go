package immutable_test

import (
	"fmt"
	"testing"

	"github.com/fishy/go-immutable"
)

func TestMapBuilder(t *testing.T) {
	orig := immutable.MapLiteralType{
		1: "a",
		2: "b",
		3: "c",
	}
	m := immutable.NewMapBuilder().Set(1, "c").Update(orig).Build()
	size := m.Len()
	if size != 3 {
		t.Errorf("Len() expected 3, got %d", size)
	}

	var key immutable.Hashable = 2
	var target interface{} = "b"
	value, ok := m.Get(key)
	if !ok {
		t.Errorf("%v should be in the map", key)
	}
	if value != target {
		t.Errorf("Get(%v) expected %v, got %v", key, target, value)
	}
	key = 0
	value, ok = m.Get(key)
	if ok {
		t.Errorf("%v should not be in the map", key)
	}
	if value != nil {
		t.Errorf("Get(%v) expected nil, got %v", key, value)
	}

	saw := immutable.NewSetBuilder()
	if err := m.Range(func(key immutable.Hashable, value interface{}) error {
		if saw.Contains(key) {
			t.Errorf("Already iterated key %v", key)
		}
		saw.Add(key)
		if value != orig[key] {
			t.Errorf("Expected %v, %v, got %v, %v", key, orig[key], key, value)
		}
		return nil
	}); err != nil {
		t.Errorf("Range() should return nil, got: %v", err)
	}
	size = saw.Len()
	if size != 3 {
		t.Errorf("Should iterated 3 key-value pairs, got %d", size)
	}
}

func TestMapString(t *testing.T) {
	orig := immutable.MapLiteralType{
		1: "a",
	}
	m := immutable.MapLiteral(orig)
	origStr := fmt.Sprintf("%v", orig)
	mapStr := fmt.Sprintf("%v", m)
	if mapStr != origStr {
		t.Errorf("Map.String() expected %q, got %q", origStr, mapStr)
	}
}

func BenchmarkMapBuilder(b *testing.B) {
	b.Run(
		"literal-5",
		func(b *testing.B) {
			b.Run(
				"baseline",
				func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						_ = map[interface{}]interface{}{
							0: 0,
							1: 1,
							2: 2,
							3: 3,
							4: 4,
						}
					}
				},
			)
			b.Run(
				"immutable",
				func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						immutable.MapLiteral(immutable.MapLiteralType{
							0: 0,
							1: 1,
							2: 2,
							3: 3,
							4: 4,
						})
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
							orig := make(immutable.MapLiteralType)
							for j := 0; j < size; j++ {
								orig[j] = j
							}
						}
					},
				)
				b.Run(
					"immutable-literal",
					func(b *testing.B) {
						for i := 0; i < b.N; i++ {
							orig := make(immutable.MapLiteralType)
							for j := 0; j < size; j++ {
								orig[j] = j
							}
							immutable.MapLiteral(orig)
						}
					},
				)
				b.Run(
					"immutable-builder",
					func(b *testing.B) {
						for i := 0; i < b.N; i++ {
							builder := immutable.NewMapBuilder()
							for j := 0; j < size; j++ {
								builder.Set(j, j)
							}
							builder.Build()
						}
					},
				)
			},
		)
	}
}

func BenchmarkMapRange(b *testing.B) {
	for _, size := range sizes {
		b.Run(
			fmt.Sprintf("%d", size),
			func(b *testing.B) {
				orig := make(immutable.MapLiteralType)
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
						m := immutable.MapLiteral(orig)
						b.ResetTimer()
						for i := 0; i < b.N; i++ {
							m.Range(func(k immutable.Hashable, v interface{}) error {
								return nil
							})
						}
					},
				)
			},
		)
	}
}
