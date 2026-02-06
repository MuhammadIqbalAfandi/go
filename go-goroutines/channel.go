//go:build ignore

package main

import (
	"fmt"
)

// <-chan T   // hanya RECEIVE (read-only)
// chan<- T   // hanya SEND (write-only)
// chan T     // SEND + RECEIVE

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}

	c <- total
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	c := make(chan int)
	defer close(c)

	go sum(nums, c)

	result := <-c

	fmt.Println("Sum:", result)
}
