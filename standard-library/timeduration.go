//go:build ignore

package main

import (
	"fmt"
	"time"
)

func main() {
	duration := 60 * time.Second
	fmt.Println(duration)
	fmt.Printf("duration: %d\n", duration)

	durationHours := 60 * time.Minute
	fmt.Println(durationHours)
	fmt.Printf("duration: %d\n", durationHours)

	durationNow := durationHours - duration
	fmt.Println(durationNow)
}