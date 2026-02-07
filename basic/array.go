//go:build ignore

package main

import "fmt"

func main() {
	// Declaring and initializing arrays in Go
	var arr [5]int

	// Short-hand declaration and initialization
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

	// Partial initialization
	arr3 := [5]int{4: 10}
	fmt.Println(arr3)

	// Array with inferred size
	arr4 := [...]int{1,2,3}
	fmt.Println(len(arr4))

	// Accessing array elements
	fmt.Println(arr[0])
	arr[0] = 42
	fmt.Println(arr[0])
	
	// Iterating over arrays
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// Array copying
	// Note: Arrays are value types in Go
	a := [3]int{1,2,3}
	b := a
	b[0] = 99
	fmt.Println(a)
	fmt.Println(b)

	// Array pointers and functions
	nums := [3]int{1,2,3}
	change(&nums)
	fmt.Println(nums)
}

func change(arr *[3]int) {
	arr[0] = 100
}