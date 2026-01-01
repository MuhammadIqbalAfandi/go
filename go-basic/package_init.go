//go:build ignore

package main

import (
	"fmt"
	"go-basic/database"
	_ "go-basic/internal"
)

func main() {
	fmt.Println(database.GetDatabase())
}