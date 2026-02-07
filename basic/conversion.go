//go:build ignore

package main

import "fmt"

func main() {
	var num1 int = 10
	var num2 float64 = 3.14

	// Implicit conversion (not allowed in Go)
	// result := num1 + num2

	// Explicit conversion
	result := float64(num1) + num2
	fmt.Println("Result:", result)
}