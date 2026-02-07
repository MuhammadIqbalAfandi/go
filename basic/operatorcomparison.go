//go:build ignore

package main

import "fmt"

func main() {
	a := 10
	b := 20
	fmt.Println(a == b)

	var c float64 = 10
	fmt.Println(float64(a) == c)

	fmt.Println("z" > "a")

	fmt.Println(true != false)
	fmt.Println(true == false)

	type User struct {
		ID int
		Name string
	}

	u1 := User{1, "Iqbal"}
	u2 := User{1, "Iqbal"}
	fmt.Println(u1 == u2)
}