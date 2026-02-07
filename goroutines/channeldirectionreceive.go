//go:build ignore

package main

import "fmt"

// <-chan T   // hanya RECEIVE (read-only)
// chan<- T   // hanya SEND (write-only)
// chan T     // SEND + RECEIVE

func consumer(in <-chan string) {
	for msg := range in {
		fmt.Println("Received:", msg)
	}
}

func main() {
	ch := make(chan string)

	go consumer(ch)
	
	ch <- "hello"
	ch <- "world"
	close(ch)
}