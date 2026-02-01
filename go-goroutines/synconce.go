//go:build ignore

package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	var wg sync.WaitGroup

	printOnce := func() {
		fmt.Println("Hanya dicetak sekali")
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Println("Goroutine", id, "memanggil once")
			once.Do(printOnce)
		}(i)
	}

	wg.Wait()
}