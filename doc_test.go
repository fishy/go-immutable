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
