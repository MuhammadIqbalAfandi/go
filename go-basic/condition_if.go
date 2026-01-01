//go:build ignore

package main

import "fmt"

func main() {
	// if statement example
	age := 21

	if age >= 18 {
		fmt.Println("You are an adult.")
	}

	// if-else statement example
	score := 65

	if score >= 80 {
		fmt.Println("A")
	} else {
		fmt.Println("B or C")
	}

	// else-if statement example
	if score >= 80 {
		fmt.Println("A")
	} else if score >= 70 {
		fmt.Println("B")
	} else if score >= 60 {
		fmt.Println("C")
	} else {
		fmt.Println("D")
	}

	// if with short statement example
	if age := 21; age >= 18 {
		fmt.Println("You are an adult.")
	}
}