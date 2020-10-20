package main

import "fmt"

func main() {
	ch1 := make(chan int, 2)
	ch1 <- 1
	ch1 <- 2
	ch2 := make(chan int, 2)
	ch2 <- 10
	ch2 <- 12

LOOP:
	for {

		select {
		case v1 := <-ch1:
			fmt.Println("V1", v1)
		case v2 := <-ch2:
			fmt.Println("V2", v2)
		default:
			break LOOP
		}
	}
}
