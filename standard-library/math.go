//go:build ignore

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Ceil(1.40))
	fmt.Println(math.Floor(1.60))
	fmt.Println(math.Round(1.60))
	fmt.Println(math.Min(1.40, 1.60))
	fmt.Println(math.Max(1.40, 1.60))
}