package immutable

import (
	"fmt"
	"strings"
)

// The 0-byte dummy value to be stored in the map.
var dummy = struct{}{}

// SetRangeFunc defines the iteration function for Set type.
//
// Whenever SetRangeFunc returns a non-nill error, the iteration will be
// stopped. The error will be returned by Range function.
type SetRangeFunc func(x Hashable) error

// Set defines the interface of an immutable set.
type Set interface {
	// Len returns the length of the set.
	Len() int
	// Contains checks whether an item is in the set.
	Contains(x Hashable) bool
	// Range iterates through the set.
	//
	// It will return the error returned by f.
	Range(f SetRangeFunc) error
}

// SetBuilder defines the interface of an immutable set builder.
type SetBuilder interface {
	Set
	// Add adds item(s) to the set.
	//
	// It should return self for chaining.
	Add(x ...Hashable) SetBuilder
	// Build builds the immutable set.
	Build() Set
}

// Make sure *set satisfies Set interface
var _ Set = (*set)(nil)

type set struct {
	m Map
}

func (s *set) Len() int {
	return s.m.Len()
}

func (s *set) Contains(x Hashable) bool {
	_, ok := s.m.Get(x)
	return ok
}

func (s *set) Range(f SetRangeFunc) error {
	return s.m.Range(func(k Hashable, _ interface{}) error {
		return f(k)
	})
}

func (s *set) String() string {
	var builder strings.Builder
	first := true
	s.Range(func(x Hashable) error {
		if first {
			builder.WriteString("[")
		} else {
			builder.WriteString(" ")
		}
		builder.WriteString(fmt.Sprintf("%v", x))
		first = false
		return nil
	})
	builder.WriteString("]")
	return builder.String()
}

// Make sure *setBuilder satisfies SetBuilder interface
var _ SetBuilder = (*setBuilder)(nil)

type setBuilder struct {
	m MapBuilder
}

func (s *setBuilder) Len() int {
	return s.m.Len()
}

func (s *setBuilder) Contains(x Hashable) bool {
	_, ok := s.m.Get(x)
	return ok
}

func (s *setBuilder) Range(f SetRangeFunc) error {
	return s.m.Range(func(k Hashable, _ interface{}) error {
		return f(k)
	})
}

func (s *setBuilder) Add(items ...Hashable) SetBuilder {
	for _, x := range items {
		s.m.Set(x, dummy)
	}
	return s
}

func (s *setBuilder) Build() Set {
	return &set{
		m: s.m.Build(),
	}
}

// NewSetBuilder creates a new SetBuilder.
func NewSetBuilder() SetBuilder {
	return &setBuilder{
		m: NewMapBuilder(),
	}
}

// SetLiteral creates an immutable set from items.
//
// It's shorthand for immutable.NewSetBuilder().Add(items...).Build().
func SetLiteral(items ...Hashable) Set {
	return NewSetBuilder().Add(items...).Build()
}
