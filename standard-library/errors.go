//go:build ignore

package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("validation error")
	NotFoundError = errors.New("not found error")
)

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "user_123" {
		return NotFoundError
	}

	return nil
}

func main() {
	err := GetById("")
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println(err)
		} else if errors.Is(err, NotFoundError) {
			fmt.Println(err)
		} else {
			fmt.Println("unknown error")
		}
	}
}