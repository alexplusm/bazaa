package main

import (
	"fmt"
)

const ZERO_ASCII_CODE = 48
const POINT_ASCII_CODE = 46

func f_power(number int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}

	return (number * f_power(number, power-1))
}

func is_point(char byte) bool {
	return char == POINT_ASCII_CODE
}

func atoi(str string) int {
	var result int
	number_rank := len(str) - 1

	for i := 0; i < len(str); i++ {
		result += int(str[i]-'0') * f_power(10, number_rank-i)
	}

	return result
}

func main() {
	ip := "123.34.54.1"

	var str string
	var ip_count_parts uint8

	for i := 0; i < len(ip); i++ {
		if !is_point(ip[i]) {
			str += string(ip[i])
		}

		if is_point(ip[i]) || ip_count_parts == 3 {
			fmt.Println(":", atoi(str))
			str = ""
			ip_count_parts += 1
		}
	}
}
