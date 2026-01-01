//go:build ignore

package main

import "fmt"

func counter() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}

func main() {
	c := counter()
	fmt.Println(c())
	fmt.Println(c())
}