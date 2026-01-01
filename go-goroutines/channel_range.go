//go:build ignore

package main

import "fmt"

func sender(c chan int) {
	for i := 1; i <= 3; i++ {
		c <- i
	}

	close(c)
}

func main() {
	c := make(chan int)

	go sender(c)

	for v := range c {
		fmt.Println(v)
	}
}