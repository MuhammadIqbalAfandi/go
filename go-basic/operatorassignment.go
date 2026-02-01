//go:build ignore

package main

import "fmt"

func main() {
	var a int = 10
	a += 5
	fmt.Println(a)

	var b int = 20
	b -= 10
	fmt.Println(b)

	var c int = 4
	c *= 3
	fmt.Println(c)

	var d int = 16
	d /= 4
	fmt.Println(d)

	var e int = 10
	e %= 3
	fmt.Println(e)
}