//go:build ignore

package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{"Toronto", "Ontario", "Canada"} 
	var address2 *Address = &address1
	address2.City = "Vancouver"
	fmt.Println("address1:", address1)
	fmt.Println("address2:", address2)

	// address2 = &Address{"Montreal", "Quebec", "Canada"}
	*address2 = Address{"Montreal", "Quebec", "Canada"}
	fmt.Println("address1:", address1)
	fmt.Println("address2:", address2)
}