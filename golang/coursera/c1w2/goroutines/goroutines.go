package main

import (
	"fmt"
	"runtime"
	"strings"
)

const (
	iterationsNum = 7
	goroutinesNum = 5
)

func doSomeWork(in int) {
	for j := 0; j <= iterationsNum; j++ {
		fmt.Println(formatWork(in, j))
		runtime.Gosched()
	}
}

func formatWork(in, j int) string {
	return fmt.Sprintln(strings.Repeat(" ", in), "*",
		strings.Repeat(" ", goroutinesNum-in), "th",
		in, "iter", j, strings.Repeat("*", j))
}

func goWithoutGoroutines() {
	for i := 0; i <= goroutinesNum; i++ {
		doSomeWork(i)
	}
}

func goGoroutines() {
	for i := 0; i <= goroutinesNum; i++ {
		go doSomeWork(i)
	}

	// wait
	fmt.Scanln()
}

func main() {
	// goWithoutGoroutines()
	goGoroutines()
}
