//go:build ignore

package main

import (
	"fmt"
	"time"
)

func download() {
	time.Sleep(2 * time.Second)
	fmt.Println("Download selesai")
}

func main() {
	go download()
	fmt.Println("program selesai")
	time.Sleep(3 * time.Second)
}
