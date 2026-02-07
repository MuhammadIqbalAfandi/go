//go:build ignore

package main

import (
	"fmt"
	"path"
)

func main() {
	strPath := "hello/world.go"
	fmt.Println(path.Dir(strPath))
	fmt.Println(path.Base(strPath))
	fmt.Println(path.Ext(strPath))
}