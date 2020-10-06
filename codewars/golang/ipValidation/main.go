package main

import (
	"fmt"
)

const zeroASCIICode = 48
const nineASCIICode = 58
const pointASCIICode = 46
const maxValidValue = 255

func isPoint(char byte) bool {
	return char == pointASCIICode
}

func isDigit(char byte) bool {
	return zeroASCIICode <= char && char <= nineASCIICode
}

func fPower(number int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}

	return (number * fPower(number, power-1))
}

func hasLeadingZero(num int, numStr string) bool {
	length := len(numStr)
	numberRank := length - 1

	for i := 0; i < numberRank; i++ {
		num /= 10
		if num == 0 {
			return true
		}
	}

	return false
}

func atoi(str string) int {
	var result int
	length := len(str)
	numberRank := length - 1

	for i := 0; i < length; i++ {
		result += int(str[i]-zeroASCIICode) * fPower(10, numberRank-i)
	}

	return result
}

func validateIPv4(ip string) bool {
	var str string
	var ipPartsCount uint8
	ipStrLength := len(ip)

	for i := 0; i < ipStrLength; i++ {
		if !(isPoint(ip[i]) || isDigit(ip[i])) {
			return false
		}

		if !isPoint(ip[i]) {
			str += string(ip[i])
		}

		if isPoint(ip[i]) || i == (ipStrLength-1) {
			if num := atoi(str); num > maxValidValue || num == -1 || hasLeadingZero(num, str) {
				return false
			}
			str = ""
			ipPartsCount++
		}
	}
	if ipPartsCount != 4 {
		return false
	}
	return true
}

func main() {
	ip := "101.230.7.80"
	res := validateIPv4(ip)
	fmt.Println("result: ", res)
}
