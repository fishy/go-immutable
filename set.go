package immutable

import (
	"fmt"
	"iter"
	"strings"
)

// The 0-byte dummy value to be stored in the map.
var dummy = struct{}{}

// SetRangeFunc defines the iteration function for Set type.
//
// Whenever SetRangeFunc returns a non-nil error, the iteration will be stopped.
// The error will be returned by Range function.
type SetRangeFunc[T comparable] func(x T) error

// Set defines the interface of an immutable set.
type Set[T comparable] interface {
	// Len returns the length of the set.
	Len() int

	// Contains checks whether an item is in the set.
	Contains(x T) bool

	// Range iterates through the set.
	//
	// It will return the error returned by f.
	Range(f SetRangeFunc[T]) error

	// All implements iter.Seq[value].
	All() iter.Seq[T]
}

// SetBuilder defines the interface of an immutable set builder.
//
// It's not guaranteed to be thread-safe and shouldn't be used concurrently.
type SetBuilder[T comparable] interface {
	Set[T]

	// Add adds item(s) to the set.
	//
	// It returns self for chaining.
	Add(x ...T) SetBuilder[T]

	// Build builds the immutable set.
	Build() Set[T]
}

// EmptySet returns an immutable empty set.
func EmptySet[T comparable]() Set[T] {
	return (*set[T])(nil)
}

type set[T comparable] struct {
	m Map[T, struct{}]
}

func (s *set[T]) Len() int {
	if s == nil {
		return 0
	}

	return s.m.Len()
}

func (s *set[T]) Contains(x T) bool {
	if s == nil {
		return false
	}

	_, ok := s.m.Load(x)
	return ok
}

func (s *set[T]) Range(f SetRangeFunc[T]) error {
	if s == nil {
		return nil
	}

	return s.m.Range(func(k T, _ struct{}) error {
		return f(k)
	})
}

func (s *set[T]) All() iter.Seq[T] {
	if s == nil {
		return func(yield func(T) bool) {}
	}

	m := s.m.All()
	return func(yield func(T) bool) {
		m(func(k T, _ struct{}) bool {
			return yield(k)
		})
	}
}

func (s *set[T]) String() string {
	var builder strings.Builder
	first := true
	s.Range(func(x T) error {
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
var _ SetBuilder[string] = (*setBuilder[string])(nil)

type setBuilder[T comparable] struct {
	m MapBuilder[T, struct{}]
}

func (s *setBuilder[T]) Len() int {
	return s.m.Len()
}

func (s *setBuilder[T]) Contains(x T) bool {
	_, ok := s.m.Load(x)
	return ok
}

func (s *setBuilder[T]) Range(f SetRangeFunc[T]) error {
	return s.m.Range(func(k T, _ struct{}) error {
		return f(k)
	})
}

func (s *setBuilder[T]) All() iter.Seq[T] {
	m := s.m.All()
	return func(yield func(T) bool) {
		m(func(k T, _ struct{}) bool {
			return yield(k)
		})
	}
}

func (s *setBuilder[T]) Add(items ...T) SetBuilder[T] {
	for _, x := range items {
		s.m.Set(x, dummy)
	}
	return s
}

func (s *setBuilder[T]) Build() Set[T] {
	return &set[T]{
		m: s.m.Build(),
	}
}

// NewSetBuilder creates a new SetBuilder.
func NewSetBuilder[T comparable]() SetBuilder[T] {
	return &setBuilder[T]{
		m: NewMapBuilder[T, struct{}](),
	}
}

// SetLiteral creates an immutable set from items.
//
// It's shorthand for immutable.NewSetBuilder[T]().Add(items...).Build().
func SetLiteral[T comparable](items ...T) Set[T] {
	return NewSetBuilder[T]().Add(items...).Build()
}
