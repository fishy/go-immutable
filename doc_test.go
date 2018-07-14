package immutable_test

import (
	"fmt"

	"github.com/fishy/go-immutable"
)

func ExampleList() {
	list := immutable.ListLiteral("a", "b", "c")
	fmt.Printf("Len: %d\n", list.Len())
	fmt.Println("Break iteration:")
	list.Range(func(i int, x interface{}) error {
		if i >= 1 {
			return immutable.ErrBreak
		}
		fmt.Printf("%d: %v\n", i, x)
		return nil
	})
	fmt.Println("Full iteration:")
	list.Range(func(i int, x interface{}) error {
		fmt.Printf("%d: %v\n", i, x)
		return nil
	})
	fmt.Printf("%%v: %v\n", list)
	// Output:
	//
	// Len: 3
	// Break iteration:
	// 0: a
	// Full iteration:
	// 0: a
	// 1: b
	// 2: c
	// %v: [a b c]
}

func ExampleMap() {
	m := immutable.MapLiteral(immutable.MapLiteralType{
		1: "a",
	})
	fmt.Printf("%%v: %v\n", m)
	m = immutable.MapLiteral(immutable.MapLiteralType{
		1: "a",
		2: "b",
		3: "c",
	})
	fmt.Printf("Len: %d\n", m.Len())
	m.Range(func(k immutable.Hashable, v interface{}) error {
		fmt.Printf("%v: %v\n", k, v)
		return nil
	})
	// Unordered Output:
	//
	// %v: map[1:a]
	// Len: 3
	// 1: a
	// 2: b
	// 3: c
}

func ExampleSet() {
	s := immutable.SetLiteral("a")
	fmt.Printf("%%v: %v\n", s)
	s = immutable.SetLiteral("a", "b", "c")
	fmt.Printf("Len: %d\n", s.Len())
	s.Range(func(x immutable.Hashable) error {
		fmt.Printf("%v\n", x)
		return nil
	})
	// Unordered Output:
	//
	// %v: [a]
	// Len: 3
	// a
	// b
	// c
}
