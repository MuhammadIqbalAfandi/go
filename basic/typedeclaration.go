//go:build ignore

package main

import "fmt"

func main() {
	// type definition
	type Age int

	// type alias
	type Text = string

	// struct
	type User struct {
			Name string
			Age  Age
	}

	// interface
	type Reader interface {
			Read() error
	}

	// function type
	type Handler func(string) error

	var a Age = 20
	var b int = int(a)

	fmt.Println(b)
}
