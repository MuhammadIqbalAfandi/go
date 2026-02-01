package go_goroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestRaceCondition(t *testing.T) {
	var counter int
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}

func TestRaceFixedMutex(t *testing.T) {
	var counter int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}

func TestRaceFixedAtomic(t *testing.T) {
	var counter int64
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}

	wg.Wait()
}

func TestRaceFixedChannel(t *testing.T) {
	counter := 0
	ch := make(chan int)
	done := make(chan struct{})

	go func() {
		for v := range ch {
			counter += v
		}
		close(done)
	}()

	for i := 0; i < 1000; i++ {
		ch <- 1
	}
	close(ch)

	<-done
	fmt.Println(counter)
}
