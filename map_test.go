package immutable_test

import (
	"fmt"
	"testing"

	"go.yhsif.com/immutable"
)

func TestMapBuilder(t *testing.T) {
	orig := map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}
	m := immutable.NewMapBuilder[int, string]().Set(1, "c").Update(orig).Build()
	size := m.Len()
	if size != 3 {
		t.Errorf("Len() expected 3, got %d", size)
	}

	key := 2
	const target = "b"
	value, ok := m.Load(key)
	if !ok {
		t.Errorf("%v should be in the map", key)
	}
	if value != target {
		t.Errorf("Load(%v) expected %v, got %v", key, target, value)
	}
	key = 0
	value, ok = m.Load(key)
	if ok {
		t.Errorf("%v should not be in the map", key)
	}
	if value != "" {
		t.Errorf("Load(%v) expected empty string, got %q", key, value)
	}

	saw := immutable.NewSetBuilder[int]()
	if err := m.Range(func(key int, value string) error {
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
	orig := map[int]string{
		1: "a",
	}
	m := immutable.MapLiteral(orig)
	origStr := fmt.Sprintf("%v", orig)
	mapStr := fmt.Sprintf("%v", m)
	if mapStr != origStr {
		t.Errorf("Map.String() expected %q, got %q", origStr, mapStr)
	}
}

func TestEmptyMap(t *testing.T) {
	if l := immutable.EmptyMap[string, any]().Len(); l != 0 {
		t.Errorf("EmptyMap.Len() expected 0, got %d", l)
	}
	if value, ok := immutable.EmptyMap[string, any]().Load("foo"); value != nil || ok {
		t.Errorf("EmptyMap.Load() expected nil, false, got %v, %v", value, ok)
	}
	if err := immutable.EmptyMap[string, any]().Range(func(k string, v any) error {
		t.Errorf("EmptyMap.Range called MapRangeFunc with %v, %v", k, v)
		return nil
	}); err != nil {
		t.Errorf("EmptyMap.Range() returned error: %v", err)
	}
}

func BenchmarkMapBuilder(b *testing.B) {
	b.Run("literal-5", func(b *testing.B) {
		b.Run("baseline", func(b *testing.B) {
			b.ReportAllocs()
			for range b.N {
				_ = map[int]int{
					0: 0,
					1: 1,
					2: 2,
					3: 3,
					4: 4,
				}
			}
		})
		b.Run("immutable", func(b *testing.B) {
			b.ReportAllocs()
			for range b.N {
				immutable.MapLiteral(map[int]int{
					0: 0,
					1: 1,
					2: 2,
					3: 3,
					4: 4,
				})
			}
		})
	})
	for _, size := range sizes {
		b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
			b.Run("baseline", func(b *testing.B) {
				b.ReportAllocs()
				for range b.N {
					orig := make(map[int]int)
					for i := 0; i < size; i++ {
						orig[i] = i
					}
				}
			})
			b.Run("immutable-literal", func(b *testing.B) {
				b.ReportAllocs()
				for range b.N {
					orig := make(map[int]int)
					for i := 0; i < size; i++ {
						orig[i] = i
					}
					immutable.MapLiteral(orig)
				}
			})
			b.Run("immutable-builder", func(b *testing.B) {
				b.ReportAllocs()
				for range b.N {
					builder := immutable.NewMapBuilder[int, int]()
					for i := 0; i < size; i++ {
						builder.Set(i, i)
					}
					builder.Build()
				}
			})
		})
	}
}

func BenchmarkMapRange(b *testing.B) {
	for _, size := range sizes {
		b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
			orig := make(map[int]int)
			for i := 0; i < size; i++ {
				orig[i] = i
			}
			b.Run("baseline", func(b *testing.B) {
				b.ReportAllocs()
				for range b.N {
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
				for range b.N {
					m.Range(func(k int, v int) error {
						return nil
					})
				}
			})
			b.Run("immutable-all", func(b *testing.B) {
				b.ReportAllocs()
				m := immutable.MapLiteral(orig)
				b.ResetTimer()
				for range b.N {
					for k, v := range m.All() {
						_ = k
						_ = v
					}
				}
			})
		})
	}
}
