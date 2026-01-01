//go:build ignore

package main

import "fmt"

func main() {
	// iterate over a range of numbers
	names := []string{"Alice", "Bob", "Charlie"}

	for index, value := range names {
		fmt.Println("Index:", index, "Value:", value)
	}

	// not using index
	for _, value := range names {
		fmt.Println("value:", value)
	}

	// iterate over a string
	text := "Golang"

	for i, ch := range text {
		fmt.Printf("%d %c\n", i, ch)
	}
}