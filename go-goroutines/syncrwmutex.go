//go:build ignore

package main

import "sync"

func main() {
	var rw sync.RWMutex
	data := 0

	var wg sync.WaitGroup

	// writer
	wg.Add(1)
	go func() {
		defer wg.Done()
		rw.Lock()
		data = 42
		rw.Unlock()
	}()

	// readers
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			rw.RLock()
			_ = data
			rw.RUnlock()
		}()
	}

	wg.Wait()
}