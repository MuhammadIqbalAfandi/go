//go:build ignore

package main

import "fmt"

func main() {
	fmt.Println("Start")
	
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	panic("there is an error")
}