//go:build ignore

package main

import (
	"errors"
	"fmt"
)

// basic function
func sayHello() {
	fmt.Println("Hello Go")
}

// function with parameters
func greet(name string) {
	fmt.Println("Hello", name)
}

// multiple parameters
func add(a, b int) int {
	return a + b
}

// return value
func square(x int) int {
	return x * x
}

// multiple return values
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// named return values
func cal(a, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return
}

// variadic function
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// function as a value
func operate(a, b int, op func(int, int) int) int {
	return op(a, b)
}

// recursive function
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	sayHello()
	greet("Iqbal")
	fmt.Println(add(1, 2))

	result, err := divide(10, 2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)

	summation, diff := cal(10, 5)
	fmt.Println(summation)
	fmt.Println(diff)

	sumResult := sum(1, 2, 3, 4, 5) 
	fmt.Println(sumResult)

	operationResult := operate(10, 5, func (x, y int) int {
		return x + y
	})
	fmt.Println(operationResult)

	// anonymous function
	func(name string) {
		fmt.Println("Hello", name)
	}("Iqbal")

	factorialResult := factorial(5)
	fmt.Println(factorialResult)
}