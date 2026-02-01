//go:build ignore

package main

import "fmt"

func main() {
	// basic for loop
	for i := 0; i < 5; i++ {
		fmt.Println("Iteration:", i)
	}

	// for loop as a while loop
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
}