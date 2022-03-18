package immutable

import (
	"fmt"
)

// ListRangeFunc defines the iteration function for List type.
//
// i will be the 0-based index and x will be the item.
//
// Whenever ListRangeFunc returns a non-nil error, the iteration will be
// stopped. The error will be returned by Range function.
type ListRangeFunc[T any] func(i int, x T) error

// List defines the interface of an immutable list.
type List[T any] interface {
	// Len returns the length of the list.
	Len() int

	// Get returns the i-th item with 0-index.
	//
	// It panics when i is out of [0, Len()-1].
	Get(i int) T

	// Range iterates through the list, in its original order.
	//
	// It will return the error returned by f.
	Range(f ListRangeFunc[T]) error

	// Reslice returns the sublist from start to end-1 index.
	//
	// Use out of range indices will cause panic.
	Reslice(start, end int) List[T]
}

// ListBuilder defines the interface of an immutable list builder.
//
// It's not guaranteed to be thread-safe and shouldn't be used concurrently.
type ListBuilder[T any] interface {
	List[T]

	// Append appends item(s) to the list.
	//
	// It returns self for chaining.
	Append(x ...T) ListBuilder[T]

	// Build builds the immutable list.
	Build() List[T]
}

// EmptyList returns an immutable empty list.
func EmptyList[T any]() List[T] {
	return (*list[T])(nil)
}

type list[T any] struct {
	list []T
}

// Make sure *list satisfies ListBuilder interface.
var _ ListBuilder[any] = (*list[any])(nil)

func (l *list[T]) Len() int {
	if l == nil {
		return 0
	}

	return len(l.list)
}

func (l *list[T]) Get(i int) T {
	return l.list[i]
}

func (l *list[T]) Range(f ListRangeFunc[T]) error {
	if l == nil {
		return nil
	}

	for i, x := range l.list {
		if err := f(i, x); err != nil {
			return err
		}
	}
	return nil
}

func (l *list[T]) Reslice(start, end int) List[T] {
	return &list[T]{list: l.list[start:end]}
}

func (l *list[T]) Append(x ...T) ListBuilder[T] {
	l.list = append(l.list, x...)
	return l
}

func (l *list[T]) Build() List[T] {
	newlist := make([]T, len(l.list))
	copy(newlist, l.list)
	return &list[T]{
		list: newlist,
	}
}

func (l *list[T]) String() string {
	return fmt.Sprintf("%v", l.list)
}

// NewListBuilder creates a ListBuilder.
func NewListBuilder[T any]() ListBuilder[T] {
	return &list[T]{}
}

// ListLiteral creates an immutable list from items.
//
// It's shorthand for immutable.NewListBuilder[T]().Append(items...).Build().
func ListLiteral[T any](items ...T) List[T] {
	return NewListBuilder[T]().Append(items...).Build()
}
