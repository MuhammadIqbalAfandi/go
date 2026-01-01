//go:build ignore

package main

import "fmt"

// <-chan T   // hanya RECEIVE (read-only)
// chan<- T   // hanya SEND (write-only)
// chan T     // SEND + RECEIVE

func producer(out chan<- int) {
	for i := 1; i <= 3; i++ {
		out <- i
	}
	close(out)
}

func main() {
	ch := make(chan int)

	go producer(ch)
	
	for v := range ch {
		fmt.Println(v)
	}
}