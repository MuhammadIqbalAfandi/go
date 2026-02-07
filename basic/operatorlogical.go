//go:build ignore

package main

import "fmt"

func main() {
	age := 20
	hasID := true

	fmt.Println(age >= 18 && hasID)

	isAdmin := false
	isOwner := true

	fmt.Println(isAdmin || isOwner)

	fmt.Println(!true)
	fmt.Println(!false)
}