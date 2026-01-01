//go:build ignore

package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup
	cond := sync.NewCond(&mu)

	ready := false

	wg.Add(2)
	
	// Goroutine 1: menunggu
	go func() {
		defer wg.Done()
		mu.Lock()
		for !ready {
			fmt.Println("Goroutine menunggu sinyal...")
			cond.Wait() // menunggu signal
		}
		fmt.Println("Goroutine bangun, lanjut eksekusi")
		mu.Unlock()
	}()

	// Goroutine 2: memberi sinyal
	go func() {
		defer wg.Done()
		mu.Lock()
		fmt.Println("mengirim sinyal...")
		ready = true
		cond.Signal() // bangunkan goroutine 1
		mu.Unlock()
	}()

	wg.Wait()
}