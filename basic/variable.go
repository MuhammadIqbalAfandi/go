//go:build ignore

package main

import "fmt"

func main() {
	var name string = "Muhammad Iqbal"
	fmt.Println("Name:", name)

	var age int = 27
	fmt.Println("Age:", age)

	// Variable grouping
	var (
		username 		= "Iqbal"
		totalCount 	= 20
		isReady   	= true
	)
	
	fmt.Println(username)
	fmt.Println(totalCount)
	fmt.Println(isReady)

	// Short Variable Declaration
	username2 := "Afandi"
	totalCount2 := 10
	isReady2 := false

	fmt.Println(username2)
	fmt.Println(totalCount2)
	fmt.Println(isReady2)
}