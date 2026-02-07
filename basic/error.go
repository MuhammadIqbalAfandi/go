//go:build ignore

package main

import (
	"errors"
	"fmt"
)

func modulus(v int, mod int) (int, error) {
	if mod == 0 {
		return 0, errors.New("zero value")
	} else {
		return v % mod, nil
	}
}

func main() {
	result, err := modulus(10, 2)
	if err == nil {
		fmt.Println("Result: ", result)
	} else {
		fmt.Println(err)
	}
}