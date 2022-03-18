package immutable_test

import (
	"fmt"

	"go.yhsif.com/immutable"
)

func ExampleList() {
	list := immutable.ListLiteral("a", "b", "c")
	fmt.Printf("Len: %d\n", list.Len())
	fmt.Println("Break iteration:")
	list.Range(func(i int, x string) error {
		if i >= 1 {
			return immutable.ErrBreak
		}
		fmt.Printf("%d: %v\n", i, x)
		return nil
	})
	fmt.Println("Full iteration:")
	list.Range(func(i int, x string) error {
		fmt.Printf("%d: %v\n", i, x)
		return nil
	})
	fmt.Printf("%%v: %v\n", list)
	fmt.Println("Reslice(1, 3):", list.Reslice(1, 3))
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
	// Reslice(1, 3): [b c]
}

func ExampleMap() {
	m := immutable.MapLiteral(map[int]string{
		1: "a",
	})
	fmt.Printf("%%v: %v\n", m)
	m = immutable.MapLiteral(map[int]string{
		1: "a",
		2: "b",
		3: "c",
	})
	fmt.Printf("Len: %d\n", m.Len())
	m.Range(func(k int, v string) error {
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
	s.Range(func(x string) error {
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

func ExampleDropOK() {
	m := immutable.MapLiteral(map[int]string{
		1: "a",
	})
	fmt.Printf("whole map: %v\n", m)
	fmt.Printf("1: %q\n", immutable.DropOK(m.Get(1))) // "a"
	fmt.Printf("2: %q\n", immutable.DropOK(m.Get(2))) // "" as this is not in the map
	// Output:
	//
	// whole map: map[1:a]
	// 1: "a"
	// 2: ""
}
