//go:build ignore

package main

import "fmt"

func main() {
	// break statement
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		
		fmt.Println(i)
	}
	
	// continue statement
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue
		}

		fmt.Println(i)
	}
}