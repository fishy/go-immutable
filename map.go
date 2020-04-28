package immutable

import (
	"fmt"
)

// MapLiteralType is the shorthand type to be used in MapLiteral.
type MapLiteralType map[Comparable]interface{}

// MapRangeFunc defines the iteration function for Map type.
//
// Whenever MapRangeFunc returns a non-nil error, the iteration will be
// stopped. The error will be returned by Range function.
type MapRangeFunc func(key Comparable, value interface{}) error

// Map defines the interface of an immutable map.
type Map interface {
	// Len returns the size of the map.
	Len() int

	// Get returns the value to the key.
	//
	// If the key is not in the map, value will be nil and ok will be false.
	Get(key Comparable) (value interface{}, ok bool)

	// Range iterates through the map.
	//
	// It will return the error returned by f.
	Range(f MapRangeFunc) error
}

// MapBuilder defines the interface of an immutable map builder.
//
// It's not guaranteed to be thread-safe and shouldn't be used concurrently.
type MapBuilder interface {
	Map

	// Set sets the key value pair to the map.
	//
	// It should return self for chaining.
	Set(key Comparable, value interface{}) MapBuilder

	// Update updates every key value pair from m to the map.
	//
	// It should return self for chaining.
	Update(m MapLiteralType) MapBuilder

	// Build builds the immutable map.
	Build() Map
}

// Make sure *immutableMap satisfies Map interface.
var _ Map = (*immutableMap)(nil)

type immutableMap struct {
	m MapLiteralType
}

func (m *immutableMap) Len() int {
	return len(m.m)
}

func (m *immutableMap) Get(key Comparable) (value interface{}, ok bool) {
	value, ok = m.m[key]
	return
}

func (m *immutableMap) Range(f MapRangeFunc) (err error) {
	for k, v := range m.m {
		err = f(k, v)
		if err != nil {
			return
		}
	}
	return
}

func (m *immutableMap) String() string {
	return fmt.Sprintf("%v", m.m)
}

// Make sure *mapBuilder satisfies MapBuilder interface.
var _ MapBuilder = (*mapBuilder)(nil)

type mapBuilder struct {
	immutableMap
}

func (mb *mapBuilder) Set(key Comparable, value interface{}) MapBuilder {
	mb.immutableMap.m[key] = value
	return mb
}

func (mb *mapBuilder) Update(incoming MapLiteralType) MapBuilder {
	for k, v := range incoming {
		mb.Set(k, v)
	}
	return mb
}

func (mb *mapBuilder) Build() Map {
	m := make(MapLiteralType)
	for k, v := range mb.immutableMap.m {
		m[k] = v
	}
	return &immutableMap{
		m: m,
	}
}

// NewMapBuilder creates a new MapBuilder.
func NewMapBuilder() MapBuilder {
	return &mapBuilder{
		immutableMap: immutableMap{
			m: make(MapLiteralType),
		},
	}
}

// MapLiteral creates an immutable map from existing map.
//
// It's shorthand for immutable.NewMapBuilder().Update(m).Build().
func MapLiteral(m MapLiteralType) Map {
	return NewMapBuilder().Update(m).Build()
}
