//go:build ignore

package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	done := make(chan struct{})

	time.AfterFunc(5 * time.Second, func() {
		fmt.Println("stoping ticker")
		ticker.Stop()
		close(done)
	})

	for {
		select {
		case t := <- ticker.C:
			fmt.Println("tick", t)
		case <- done:
			fmt.Println("done")
			return
		}
	}
}