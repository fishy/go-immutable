package immutable_test

import (
	"fmt"

	"go.yhsif.com/immutable"
)

func ExampleList() {
	list := immutable.ListLiteral("a", "b", "c")
	fmt.Println("Len:", list.Len())
	fmt.Println("list.Get(1):", list.Get(1))
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
	// list.Get(1): b
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
	fmt.Println("Len:", m.Len())
	fmt.Println("m.Get(1):", m.Get(1))
	fmt.Println("range:")
	m.Range(func(k int, v string) error {
		fmt.Printf("%v: %v\n", k, v)
		return nil
	})
	// Unordered Output:
	//
	// %v: map[1:a]
	// Len: 3
	// m.Get(1): a
	// range:
	// 1: a
	// 2: b
	// 3: c
}

func ExampleSet() {
	s := immutable.SetLiteral("a")
	fmt.Printf("%%v: %v\n", s)
	s = immutable.SetLiteral("a", "b", "c")
	fmt.Println("Len:", s.Len())
	fmt.Println(`s.Contains("a"):`, s.Contains("a"))
	fmt.Println(`s.Contains("d"):`, s.Contains("d"))
	fmt.Println("range:")
	s.Range(func(x string) error {
		fmt.Printf("%v\n", x)
		return nil
	})
	// Unordered Output:
	//
	// %v: [a]
	// Len: 3
	// s.Contains("a"): true
	// s.Contains("d"): false
	// range:
	// a
	// b
	// c
}
