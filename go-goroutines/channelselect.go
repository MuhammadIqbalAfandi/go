//go:build ignore

package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "data dari ch1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "data dari ch2"
	}()

	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case msg := <- ch2:
		fmt.Println(msg)
	}
}