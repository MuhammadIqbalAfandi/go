//go:build ignore

package main

import "fmt"

// struct definition
type User struct {
	ID string
	Name string
	Age int
}

// method for User struct
func (u User) Greet() {
	fmt.Println("Hello", u.Name)
}

// struct composition
type Address struct {
	City string
}

type Employee struct {
	User
	Address
	Salary int
}

func main() {
	u := User {
		ID: "U01",
		Name: "Alice",
		Age: 30,
	}

	fmt.Println(u)
	fmt.Println(u.Name)

	u.Greet()

	e := Employee {
		User: User {
			ID: "E01",
			Name: "Bob",
			Age: 25,
		},
		Address: Address {
			City: "Jakarta",
		},
		Salary: 50000,
	}

	fmt.Println(e)
	fmt.Println(e.User.Name)
	fmt.Println(e.Salary)

	// Anonymous struct
	blog := struct {
		Title string
		Page string
	} {
		Title: "My Blog",
		Page: "Home",
	}

	fmt.Println(blog)
	fmt.Println(blog.Title)
}