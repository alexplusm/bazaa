package main

import "fmt"

func main() {
	ch1 := make(chan int)

	go func(out chan<- int) {
		for i := 0; i < 10; i++ {
			fmt.Println("before:", i)
			out <- i
			fmt.Println("after:", i)
		}
		close(out)
		fmt.Println("generation end")
	}(ch1)

	for i := range ch1 {
		fmt.Println("VALUE", i)
	}
}
