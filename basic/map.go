//go:build ignore

package main

import "fmt"

func main() {
	// make a map
	score := make(map[string]int)
	score["iqbal"] = 90
	score["grace"] = 85
	
	fmt.Println(score)
	fmt.Println(score["grace"])
	
	// declare and initialize a map
	rank := map[string]int{
		"iqbal": 1,
		"grace": 2,
	}

	fmt.Println(rank)

	// var map and make
	var age map[string]int
	age = make(map[string]int)
	age["iqbal"] = 21
	fmt.Println(age)

	// delete from map
	delete(rank, "iqbal")
	fmt.Println(rank)

	// iterate over map
	for k, v := range score {
		fmt.Println(k, v)
	}
}