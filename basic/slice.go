//go:build ignore

package main

import "fmt"

func main() {
	// Declaring and initializing slices in Go
	var nums []int
	fmt.Println(nums == nil)

	// Short-hand declaration and initialization
	nums1 := []int{1,2,3}
	fmt.Println(nums1)

	// Accessing and modifying slice elements
	nums4 := []int{10, 20, 30}
	fmt.Println(nums4)
	nums4[0] = 100
	fmt.Println(nums4)
	fmt.Println(nums4[0])

	// make slice and append
	nums3 := make([]int, 2,5)
	fmt.Println(len(nums3))
	fmt.Println(cap(nums3))
	nums3 = append(nums3, 3)
	nums3 = append(nums3, 4,5,6)
	fmt.Println(nums3)

	// reference type behavior
	// Note: Slices are reference types in Go
	a := []int{1,2,3}
	b := a
	b[0] = 99
	fmt.Println(a)
	fmt.Println(b)

	// Copying slices
	src := []int{1,2,3}
	dest := make([]int, 2)
	n := copy(dest, src)
	fmt.Println(n)
	fmt.Println(dest)

	// Slices
	names := []string{"Alice", "Bob", "Charlie", "Diana"}
	fmt.Println(names)

	c := names[1:3]
	fmt.Println(c)

	d := names[:2]
	fmt.Println(d)

	e := names[2:]
	fmt.Println(e)

	f := names[:]
	fmt.Println(f)

	// Lopping through slices
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
}