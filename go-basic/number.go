//go:build ignore

package main

import "fmt"

func main() {
	var i8 int8 = -8
	var i int = -100

	fmt.Println("==signed integer==")
	fmt.Println("i8 = ", i8)
	fmt.Println("i = ", i)

	var u64 uint64 = 64
	var u uint = 100

	fmt.Println("==unsigned integer==")
	fmt.Println("u64 = ", u64)
	fmt.Println("u = ", u)

	var b byte = 255
	var rune = 'A'

	fmt.Println("==byte and rune==")
	fmt.Println("b = ", b)
	fmt.Println("rune = ", rune)

	var f32 float32 = 3.14
	var f64 float64 = 3.14159265359

	fmt.Println("==floating point==")
	fmt.Println("f32 = ", f32)
	fmt.Println("f64 = ", f64)
}
