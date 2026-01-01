//go:build ignore

package main

import "fmt"

func main() {
	const defaultTimeOut = 5

	const (
		statusOk     = 200
		statusError  = 500
	)

	fmt.Println(defaultTimeOut)
	fmt.Println(statusOk)
	fmt.Println(statusError)
}