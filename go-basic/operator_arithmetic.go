//go:build ignore

package main

import "fmt"

func main() {
	a := 10;
	b := 5;
	fmt.Println(a + b)

	s := "Go" + "Lang"
	fmt.Println(s)

	fmt.Println(7 / 2)
	fmt.Println(7.0 / 2)
	fmt.Println(float64(7) / 2)

	fmt.Println(10 % 3)

	x := 2
	x++
	fmt.Println(x)
}