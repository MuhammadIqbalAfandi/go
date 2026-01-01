//go:build ignore

package main

import "fmt"

func change(x *int) {
	*x = 100
}

type Address struct {
	City, Province, Country string
}

func changeCountry(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	a := 10
	change(&a)
	fmt.Println(a)

	address := Address{}
	changeCountry(&address)
	fmt.Println(address)
}