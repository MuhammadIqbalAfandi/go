//go:build ignore

package main

import (
	"fmt"
	"sync"
)

type MyStruct struct {
	Value int
}

func main() {
	var pool = sync.Pool {
		New: func() interface{} {
			fmt.Println("Membuat object baru")
			return &MyStruct{}
		},
	}

	obj1 := pool.Get().(*MyStruct)
	obj1.Value = 42
	fmt.Println("obj1:", obj1.Value)

	pool.Put(obj1)

	obj2 := pool.Get().(*MyStruct)
	fmt.Println("obj2:", obj2.Value)

	obj3 := pool.Get().(*MyStruct)
	fmt.Println("obj3:", obj3.Value)
}