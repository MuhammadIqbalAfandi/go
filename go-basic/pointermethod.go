//go:build ignore

package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	m := Man{"Iqbal"}
	m.Married()
	fmt.Println(m.Name)
}