//go:build ignore

package main

// func main() {
// 	c := make(chan int, 2)

// 	c <- 1
// 	c <- 2

// 	fmt.Println(<- c)
// 	fmt.Println(<- c)
// 	fmt.Println(cap(c))
// 	fmt.Println(len(c))
// }

func worker(c chan int) {
	c <- 10
	c <- 20
}

func main() {
	c := make(chan int)

	go worker(c)

	println(<-c)
	println(<-c)
}