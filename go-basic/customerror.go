//go:build ignore

package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{ "ID cannot be empty"}
	} 

	if id != "user-123" {
		return &notFoundError{ "Data not found"}
	}

	return nil
}

func main() {
	err := SaveData("", nil)
	if err != nil {
		switch e := err.(type) {
		case *validationError:
			fmt.Println("Validation Error:", e.Message)
		case *notFoundError:
			fmt.Println("Not Found Error:", e.Message)
		default:
			fmt.Println("Unknown Error:", e.Error())
		}
	} else {
		fmt.Println("Success")
	}
}