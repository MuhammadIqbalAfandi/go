//go:build ignore

package main

import "fmt"

type Greeter interface {
	Greet()
}

type User struct{}

func (User) Greet() {}

// type assertion
func process(v interface{}) {
    switch val := v.(type) {
    case string:
        fmt.Println("string:", val)
    case int:
        fmt.Println("int:", val)
    default:
        fmt.Println("unknown")
    }
}

func main() {
	// basic type assertion
	var x interface{} = "hello"
	s, ok := x.(string)
	if ok {
		fmt.Println(s)
	}

	// assertion with interface
	var g interface{} = User{}
	if _, ok := g.(User); ok {
		fmt.Println("This user")
	}

	process("hello")
	process(27)
}