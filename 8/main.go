package main

import "fmt"

const minValueAbs uint64 = 2147483648
const maxValueAbs uint64 = 2147483647

func myAtoi(str string) int {
	if str == "" || str == "-" {
		return 0
	}
	index := 0
	minus := false
	var value uint64
	for index < len(str) && str[index] == ' ' {
		index++
	}
	if index >= len(str) {
		return 0
	}
	if index < len(str) && str[index] == '-' {
		minus = true
		index++
	} else if index < len(str) && str[index] == '+' {
		index++
	}
	if index >= len(str) {
		return 0
	}
	if !('0' <= str[index] && str[index] <= '9') {
		return 0
	}
	value = uint64(str[index] - '0')
	index++
	for ; index < len(str) && '0' <= str[index] && str[index] <= '9'; index++ {
		value *= 10
		value += uint64(str[index] - '0')
		if value > minValueAbs && minus {
			return -2147483648
		} else if value > maxValueAbs && !minus {
			return 2147483647
		}
	}
	if minus {
		return -int(value)
	}
	return int(value)
}

func main() {
	fmt.Println(myAtoi("-91283472332"))
}
