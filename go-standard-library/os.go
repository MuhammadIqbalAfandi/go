//go:build ignore

package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Args provides access to raw command-line arguments.
	input := os.Args
	for i, v := range input {
		fmt.Println("Index:", i, "Value:", v)
	}

	// os.Hostname gets the hostname of the system.
	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("Error", err.Error())
	}
}