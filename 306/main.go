package main

import (
	"math/big"
	"strings"
)

func isAdditiveNumberWithLastResult(str string, last string) bool {
	for i := 1; i < len(str); i++ {
		if len(str[:i]) >= 2 && str[0] == '0' {
			break
		}
		addedResult := big.NewInt(0)
		addedResult.SetString(last, 10)
		thisSlice := big.NewInt(0)
		thisSlice.SetString(str[:i], 10)
		addedResult.Add(thisSlice, addedResult)
		if strings.HasPrefix(str[i:], addedResult.String()) {
			if str[i:] == addedResult.String() {
				return true
			}
			if isAdditiveNumberWithLastResult(str[i:], str[:i]) {
				return true
			}
		}
	}
	return false
}

func isAdditiveNumber(num string) bool {
	for i := 1; i < len(num); i++ {
		if len(num[:i]) >= 2 && num[0] == '0' {
			return false
		}
		if isAdditiveNumberWithLastResult(num[i:], num[:i]) {
			return true
		}
	}
	return false
}

func main() {
	println(isAdditiveNumber("0235813"))
}
