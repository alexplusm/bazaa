package main

import (
	"fmt"
	"strconv"
	"strings"
)

func highAndLow(str string) string {
	chars := strings.Split(str, " ")
	num, err := strconv.Atoi(chars[0])

	if err != nil {
		return ""
	}

	high := num
	low := num

	for i := 0; i < len(chars); i++ {
		num, err := strconv.Atoi(chars[i])

		if err != nil {
			return ""
		}

		if num > high {
			high = num
		}

		if num < low {
			low = num
		}
	}

	return strconv.Itoa(high) + " " + strconv.Itoa(low)
}

func main() {
	str := "1 2 3 4 5"
	res := "5 1"
	myRes := highAndLow(str)

	fmt.Println(myRes == res)
}
