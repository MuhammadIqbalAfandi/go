//go:build ignore

package main

import (
	"fmt"
	"strconv"
)

func main() {
	resultBool, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(resultBool)
	}

	resultInt, err := strconv.Atoi("1000")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(resultInt)
	}

	stringInt := strconv.Itoa(100)
	fmt.Println(stringInt)

	binaryStr := strconv.FormatInt(100, 2)
	fmt.Println(binaryStr)
}