package immutable

import (
	"fmt"
	"sync"
)

// ListRangeFunc defines the iteration function for List type.
//
// i will be the 0-based index and x will be the item.
//
// Whenever ListRangeFunc returns a non-nill error, the iteration will be
// stopped. The error will be returned by Range function.
type ListRangeFunc func(i int, x interface{}) error

// List defines the interface of an immutable list.
type List interface {
	// Len returns the length of the list.
	Len() int
	// Range iterates through the list, in its original order.
	//
	// It will return the error returned by f.
	Range(f ListRangeFunc) error
}

// ListBuilder defines the interface of an immutable list builder.
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

	// needLock should be true if it's used as builder, and false if it's used as
	// immutable list.
	needLock bool
	lock     sync.RWMutex
}

// Make sure *list satisfies List and ListBuilder interfaces.
var (
	_ List        = (*list)(nil)
	_ ListBuilder = (*list)(nil)
)

func (l *list) Len() int {
	if l.needLock {
		l.lock.RLock()
		defer l.lock.RUnlock()
	}
	return len(l.list)
}

func (l *list) Range(f ListRangeFunc) error {
	if l.needLock {
		l.lock.RLock()
		defer l.lock.RUnlock()
	}
	for i, x := range l.list {
		if err := f(i, x); err != nil {
			return err
		}
	}
	return nil
}

func (l *list) Append(x ...interface{}) ListBuilder {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.list = append(l.list, x...)
	return l
}

func (l *list) Build() List {
	l.lock.RLock()
	defer l.lock.RUnlock()
	newlist := make([]interface{}, len(l.list))
	copy(newlist, l.list)
	return &list{
		list:     newlist,
		needLock: false,
	}
}

func (l *list) String() string {
	if l.needLock {
		l.lock.RLock()
		defer l.lock.RUnlock()
	}
	return fmt.Sprintf("%v", l.list)
}

// NewListBuilder creates a ListBuilder.
func NewListBuilder() ListBuilder {
	return &list{
		needLock: true,
	}
}

// ListLiteral creates an immutable list from items.
//
// It's shorthand for immutable.NewListBuilder().Append(items...).Build().
func ListLiteral(items ...interface{}) List {
	return NewListBuilder().Append(items...).Build()
}
