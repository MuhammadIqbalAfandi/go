//go:build ignore

package main

import "fmt"

// Define an interface
type Greeter interface {
	Greet()
}

type User struct {
	Name string
}

func (u User) Greet() {
	fmt.Println("Hello", u.Name)
}

// Interface as parameter
func SayHello(g Greeter) {
	g.Greet()
}

func main() {
	var g Greeter
	g = User {Name: "Alice"}
	g.Greet()

	u := User{Name: "Bob"}
	SayHello(u)

	// Empty interface
	// var x interface {}
	var x any
	x = 10
	x = "Hello"
	fmt.Println(x)
}