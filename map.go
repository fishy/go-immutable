package immutable

import (
	"fmt"
)

// MapRangeFunc defines the iteration function for Map type.
//
// Whenever MapRangeFunc returns a non-nil error, the iteration will be
// stopped. The error will be returned by Range function.
type MapRangeFunc[K comparable, V any] func(key K, value V) error

// Map defines the interface of an immutable map.
type Map[K comparable, V any] interface {
	// Len returns the size of the map.
	Len() int

	// Get returns the value to the key.
	//
	// If the key is not in the map, value will be zero and ok will be false.
	Get(key K) (value V, ok bool)

	// Range iterates through the map.
	//
	// It will return the error returned by f.
	Range(f MapRangeFunc[K, V]) error
}

// MapBuilder defines the interface of an immutable map builder.
//
// It's not guaranteed to be thread-safe and shouldn't be used concurrently.
type MapBuilder[K comparable, V any] interface {
	Map[K, V]

	// Set sets the key value pair to the map.
	//
	// It returns self for chaining.
	Set(key K, value V) MapBuilder[K, V]

	// Update updates every key value pair from m to the map.
	//
	// It returns self for chaining.
	Update(m map[K]V) MapBuilder[K, V]

	// Build builds the immutable map.
	Build() Map[K, V]
}

// EmptyMap returns an immutable empty map.
func EmptyMap[K comparable, V any]() Map[K, V] {
	return (*immutableMap[K, V])(nil)
}

type immutableMap[K comparable, V any] struct {
	m map[K]V
}

func (m *immutableMap[K, V]) Len() int {
	if m == nil {
		return 0
	}

	return len(m.m)
}

func (m *immutableMap[K, V]) Get(key K) (value V, ok bool) {
	if m == nil {
		return
	}

	value, ok = m.m[key]
	return
}

func (m *immutableMap[K, V]) Range(f MapRangeFunc[K, V]) (err error) {
	if m == nil {
		return
	}

	for k, v := range m.m {
		err = f(k, v)
		if err != nil {
			return
		}
	}
	return
}

func (m *immutableMap[K, V]) String() string {
	return fmt.Sprintf("%v", m.m)
}

// Make sure *mapBuilder satisfies MapBuilder interface.
var _ MapBuilder[int, any] = (*mapBuilder[int, any])(nil)

type mapBuilder[K comparable, V any] struct {
	immutableMap[K, V]
}

func (mb *mapBuilder[K, V]) Set(key K, value V) MapBuilder[K, V] {
	mb.immutableMap.m[key] = value
	return mb
}

func (mb *mapBuilder[K, V]) Update(incoming map[K]V) MapBuilder[K, V] {
	for k, v := range incoming {
		mb.Set(k, v)
	}
	return mb
}

func (mb *mapBuilder[K, V]) Build() Map[K, V] {
	m := make(map[K]V)
	for k, v := range mb.immutableMap.m {
		m[k] = v
	}
	return &immutableMap[K, V]{
		m: m,
	}
}

// NewMapBuilder creates a new MapBuilder.
func NewMapBuilder[K comparable, V any]() MapBuilder[K, V] {
	return &mapBuilder[K, V]{
		immutableMap: immutableMap[K, V]{
			m: make(map[K]V),
		},
	}
}

// MapLiteral creates an immutable map from existing map.
//
// It's shorthand for immutable.NewMapBuilder[K, V]().Update(m).Build().
func MapLiteral[K comparable, V any](m map[K]V) Map[K, V] {
	return NewMapBuilder[K, V]().Update(m).Build()
}
