//go:build ignore

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOMAXPROCS(-1))
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

	runtime.GOMAXPROCS(20)
	fmt.Println(runtime.GOMAXPROCS(-1))
}