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
type ListRangeFunc func(i int, x interface{}) error

// List defines the interface of an immutable list.
type List interface {
	// Len returns the length of the list.
	Len() int

	// Get returns the i-th item with 0-index.
	//
	// It panics when i is out of [0, Len()-1].
	Get(i int) interface{}

	// Range iterates through the list, in its original order.
	//
	// It will return the error returned by f.
	Range(f ListRangeFunc) error

	// Reslice returns the sublist from start to end-1 index.
	//
	// Use out of range indices will cause panic.
	Reslice(start, end int) List
}

// ListBuilder defines the interface of an immutable list builder.
//
// It's not guaranteed to be thread-safe and shouldn't be used concurrently.
type ListBuilder interface {
	List

	// Append appends item(s) to the list.
	//
	// It should return self for chaining.
	Append(x ...interface{}) ListBuilder

	// Build builds the immutable list.
	Build() List
}

type list struct {
	list []interface{}
}

// Make sure *list satisfies List and ListBuilder interfaces.
var (
	_ List        = (*list)(nil)
	_ ListBuilder = (*list)(nil)
)

func (l *list) Len() int {
	return len(l.list)
}

func (l *list) Get(i int) interface{} {
	return l.list[i]
}

func (l *list) Range(f ListRangeFunc) error {
	for i, x := range l.list {
		if err := f(i, x); err != nil {
			return err
		}
	}
	return nil
}

func (l *list) Reslice(start, end int) List {
	return &list{list: l.list[start:end]}
}

func (l *list) Append(x ...interface{}) ListBuilder {
	l.list = append(l.list, x...)
	return l
}

func (l *list) Build() List {
	newlist := make([]interface{}, len(l.list))
	copy(newlist, l.list)
	return &list{
		list: newlist,
	}
}

func (l *list) String() string {
	return fmt.Sprintf("%v", l.list)
}

// NewListBuilder creates a ListBuilder.
func NewListBuilder() ListBuilder {
	return &list{}
}

// ListLiteral creates an immutable list from items.
//
// It's shorthand for immutable.NewListBuilder().Append(items...).Build().
func ListLiteral(items ...interface{}) List {
	return NewListBuilder().Append(items...).Build()
}
