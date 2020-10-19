package main

import "fmt"

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int)

	ch1 <- 2
	select {
	case val := <-ch1:
		fmt.Println("chan 1", val)
	case ch2 <- 1:
		fmt.Println("put value to chan 2")
	default:
		fmt.Println("default")
	}
}
