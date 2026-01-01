//go:build ignore

package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map

	// Store value
	m.Store("a", 10)
	m.Store("b", 20)

	// Get value
	if v, ok := m.Load("a"); ok {
		fmt.Println("a =", v)
	}

	// Iteration
	m.Range(func(key, value interface{}) bool {
		fmt.Println(key, "=", value)
		return true
	})

	// LoadOrStore
	actual, loaded := m.LoadOrStore("a", 100)
	fmt.Println("actual:", actual, "loaded:", loaded)

	// Delete key
	m.Delete("b")

	fmt.Println(m.LoadOrStore("b", 100))
}