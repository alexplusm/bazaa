package main

import "fmt"

func findOdd(seq []int) int {
	numMap := map[int]int{}

	for _, num := range seq {
		numMap[num]++
	}

	for key, value := range numMap {
		if value%2 == 1 {
			return key
		}
	}

	return -1
}

func main() {
	nums := []int{1, 2, 1, 2, 5}
	res := findOdd(nums)
	fmt.Println(res)
}
