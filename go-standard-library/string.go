//go:build ignore

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Muhammad Iqbal Afandi", "Iqbal"))
	fmt.Println(strings.Split("Muhammad Iqbal Afandi", " "))
	fmt.Println(strings.ToLower("Muhammad Iqbal Afandi"))
	fmt.Println(strings.ToUpper("Muhammad Iqbal Afandi"))
	fmt.Println(strings.Trim("!!! Muhammad Iqbal Afandi !!!", "! "))
	fmt.Println(strings.ReplaceAll("Muhammad Iqbal", "Iqbal", "Ummar"))
}