package main

import (
	"fmt"
	"sync"
)

func withoutMutex() {
	var counters = map[int]int{}

	for i := 0; i < 5; i++ {
		go func(counters map[int]int, th int) {
			for j := 0; j < 5; j++ {
				counters[th*10+j]++
			}
		}(counters, i)
	}
	fmt.Scanln()
	fmt.Println(counters)
}

func withMutex() {
	var counters = map[int]int{}
	mu := &sync.Mutex{}

	for i := 0; i < 5; i++ {
		go func(counters map[int]int, th int, mu *sync.Mutex) {
			for j := 0; j < 5; j++ {
				mu.Lock()
				counters[th*10+j]++
				mu.Unlock()
			}
		}(counters, i, mu)
	}
	fmt.Scanln()
	mu.Lock()
	fmt.Println(counters)
	mu.Unlock()
}

func main() {
	// withoutMutex()
	withMutex()
}
