package main

import (
	"fmt";
	"strings"
)

func duplicateCount(s1 string) int {
	var dublicatesCount int
	charMap := map[byte]int{}
	lowerStr := strings.ToLower(s1)
	strBytes := []byte(lowerStr)

	for _, value := range strBytes {
		charMap[value]++
	}

	for _, value := range charMap {
		if value > 1 {
			dublicatesCount++
		}
	}

	return dublicatesCount
}

func main() {
	dublicatesCount := duplicateCount("12345678901234567890	")

	fmt.Println(dublicatesCount)
}
