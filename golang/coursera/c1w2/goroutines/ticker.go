package main

import (
	"fmt"
	"time"
)

func tickerExample() {
	ticker := time.NewTicker(time.Second)
	i := 0

	for tickerTime := range ticker.C {
		i++

		if i > 5 {
			ticker.Stop()
			break
		}
		fmt.Println(tickerTime)
	}
}

func sayHello() {
	fmt.Println("async hello")
}

func afterFuncExample() *time.Timer {
	return time.AfterFunc(time.Second, sayHello)
}

func main() {
	// go ticker Example()
	// afterFuncExample()
	// fmt.Scanln()
	// ----

	timer := afterFuncExample()

	fmt.Scanln()

	timer.Stop()

	fmt.Scanln()
}
