package immutable_test

import (
	"fmt"

	immutable "github.com/fishy/go-immutable"
)

type ImmutableIntStringMap struct {
	m immutable.Map
}

func (m ImmutableIntStringMap) Len() int {
	return m.m.Len()
}

func (m ImmutableIntStringMap) Get(key int) (value string, ok bool) {
	var v immutable.Comparable
	v, ok = m.m.Get(key)
	if ok {
		value = v.(string)
	}
	return
}

func (m ImmutableIntStringMap) Range(f func(key int, value string) error) error {
	return m.m.Range(func(k immutable.Comparable, v interface{}) error {
		return f(k.(int), v.(string))
	})
}

// This example demonstrates how to wrap immutable.Map into a stronger type map
// (we use int -> string as an example).
func ExampleMap_wrapped() {
	m := ImmutableIntStringMap{
		m: immutable.MapLiteral(immutable.MapLiteralType{
			1: "a",
			2: "b",
			3: "c",
		}),
	}
	fmt.Printf("Len: %d\n", m.Len())
	m.Range(func(k int, v string) error {
		fmt.Printf("%v: %v\n", k, v)
		return nil
	})
	// Unordered Output:
	//
	// Len: 3
	// 1: a
	// 2: b
	// 3: c
}
