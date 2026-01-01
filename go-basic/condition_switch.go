//go:build ignore

package main

import "fmt"

func main() {
	// basic switch statement
	day := 3

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Invalid day")
	}

	// switch with multiple cases
	switch day {
	case 1,2,3,4,5: 
		fmt.Println("Weekday")
	case 6,7:
		fmt.Println("Weekend")
	}

	// without expression
	score := 75

	switch {
	case score >= 80:
		fmt.Println("A")
	case score >= 70:
		fmt.Println("B")
	case score >= 60:
		fmt.Println("C")
	default:
		fmt.Println("D") 
	}
}