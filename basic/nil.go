//go:build ignore

package main

import "fmt"

type User struct {
	Name string
}

type Greeter interface {
	Greet()
}

func (u *User) Greet() {
	if  u == nil {
		fmt.Println("user is nill")
		return
	}

	fmt.Println(u.Name)
}

func main() {
	var u *User
	u.Greet()
}