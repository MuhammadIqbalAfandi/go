//go:build ignore

package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			fmt.Println("heartbeat")
		}
	}()

	time.Sleep(5 * time.Second)
}