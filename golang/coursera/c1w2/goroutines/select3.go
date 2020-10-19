package main

import "fmt"

func main() {
	cancelCh := make(chan struct{})
	dataCh := make(chan int)

	go func(cancelCh chan struct{}, dataCh chan int) {
		val := 0
		for {
			select {
			// case data:= <-dataCh:
			case <-cancelCh:
				return
			case dataCh <- val:
				val++
			}
		}
	}(cancelCh, dataCh)

	for curVal := range dataCh {
		fmt.Println("Cur Val", curVal)
		if curVal > 10 {
			fmt.Println("Stop")
			cancelCh <- struct{}{}
			break
		}
	}
}
