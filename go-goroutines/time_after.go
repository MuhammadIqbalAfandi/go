//go:build ignore

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Waiting...")
	<-time.After(2 * time.Second)
	fmt.Println("done")
}