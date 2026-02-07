//go:build ignore

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	var wg sync.WaitGroup

	ready := false
	workers := 3

	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func(id int) {
			defer wg.Done()

			mu.Lock()
			for !ready {
				fmt.Println("Worker", id, "menunggu")
				cond.Wait()
			}
			fmt.Println("Worker", id, "jalan")
			mu.Unlock()
		}(i)
	}

	time.Sleep(time.Second)

	mu.Lock()
	fmt.Println("Controller: ready = true (Signal)")
	ready = true
	cond.Broadcast()
	mu.Unlock()

	wg.Wait()
}
