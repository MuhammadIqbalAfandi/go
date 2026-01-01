//go:build ignore

package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Rob", "Pike", "Ken", "Robert"}
	values := []int{100, 90, 80, 85}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "Pike"))
}