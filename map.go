package immutable

import (
	"fmt"
	"sync"
)

// MapLiteralType is the shorthand type to be used in MapLiteral.
type MapLiteralType map[Hashable]interface{}

// MapRangeFunc defines the iteration function for Map type.
//
// Whenever MapRangeFunc returns a non-nil error, the iteration will be
// stopped. The error will be returned by Range function.
type MapRangeFunc func(key Hashable, value interface{}) error

// Map defines the interface of an immutable map.
type Map interface {
	// Len returns the size of the map.
	Len() int
	// Get returns the value to the key.
	//
	// If the key is not in the map, value will be nil and ok will be false.
	Get(key Hashable) (value interface{}, ok bool)
	// Range iterates through the map.
	//
	// It will return the error returned by f.
	Range(f MapRangeFunc) error
}

// MapBuilder defines the interface of an immutable map builder.
type MapBuilder interface {
	Map
	// Set sets the key value pair to the map.
	//
	// It should return self for chaining.
	Set(key Hashable, value interface{}) MapBuilder
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

func (m *immutableMap) Get(key Hashable) (value interface{}, ok bool) {
	v, ok := m.m[key]
	return v, ok
}

func (m *immutableMap) Range(f MapRangeFunc) error {
	for k, v := range m.m {
		if err := f(k, v); err != nil {
			return err
		}
	}
	return nil
}

func (m *immutableMap) String() string {
	return fmt.Sprintf("%v", m.m)
}

// Make sure *mapBuilder satisfies MapBuilder interface.
var _ MapBuilder = (*mapBuilder)(nil)

type mapBuilder struct {
	m sync.Map
}

func (m *mapBuilder) Len() int {
	n := 0
	m.m.Range(func(k, v interface{}) bool {
		n++
		return true
	})
	return n
}

func (m *mapBuilder) Get(key Hashable) (value interface{}, ok bool) {
	return m.m.Load(key)
}

func (m *mapBuilder) Range(f MapRangeFunc) (retErr error) {
	m.m.Range(func(k, v interface{}) bool {
		if err := f(Hashable(k), v); err != nil {
			retErr = err
			return false
		}
		return true
	})
	return
}

func (m *mapBuilder) Set(key Hashable, value interface{}) MapBuilder {
	m.m.Store(key, value)
	return m
}

func (m *mapBuilder) Update(incoming MapLiteralType) MapBuilder {
	for k, v := range incoming {
		m.Set(k, v)
	}
	return m
}

func (m *mapBuilder) Build() Map {
	newMap := make(MapLiteralType)
	m.m.Range(func(k, v interface{}) bool {
		newMap[k] = v
		return true
	})
	return &immutableMap{
		m: newMap,
	}
}

// NewMapBuilder creates a new MapBuilder.
func NewMapBuilder() MapBuilder {
	return &mapBuilder{}
}

// MapLiteral creates an immutable map from existing map.
//
// It's shorthand for immutable.NewMapBuilder().Update(m).Build().
func MapLiteral(m MapLiteralType) Map {
	return NewMapBuilder().Update(m).Build()
}
