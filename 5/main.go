package main

import (
	"fmt"
)

const seperateCharacter = byte(127)

func generateSeparatedString(s []byte) []byte {
	result := []byte{'$', seperateCharacter}
	for _, ch := range s {
		result = append(result, ch)
		result = append(result, seperateCharacter)
	}
	return result
}

func minInt(x int, y int) int {
	if x > y {
		return y
	}
	return x
}

func maxInt(x int, y int) int {
	if x < y {
		return y
	}
	return x
}

func manacher(s []byte) []byte {
	seperated := generateSeparatedString(s)
	var palindromeLengthToRight []int
	for range seperated {
		palindromeLengthToRight = append(palindromeLengthToRight, 0)
	}
	maxPalindromeLengthToRightValue := 0
	maxPalindromeLengthToRightIndex := 0
	resultStartIndex := 0
	resultLength := 0
	for i := range seperated {
		if i < maxPalindromeLengthToRightValue {
			palindromeLengthToRight[i] = minInt(
				maxPalindromeLengthToRightValue-i,
				palindromeLengthToRight[2*maxPalindromeLengthToRightIndex-i])
		} else {
			palindromeLengthToRight[i] = 1
		}
		for i >= palindromeLengthToRight[i] && i+palindromeLengthToRight[i] < len(seperated) &&
			seperated[i-palindromeLengthToRight[i]] == seperated[i+palindromeLengthToRight[i]] {
			palindromeLengthToRight[i]++
		}
		if palindromeLengthToRight[i]+i > maxPalindromeLengthToRightValue {
			maxPalindromeLengthToRightValue = palindromeLengthToRight[i] + i
			maxPalindromeLengthToRightIndex = i
		}
		if resultLength < int(palindromeLengthToRight[i]-1) {
			resultStartIndex = (i - int(palindromeLengthToRight[i])) / 2
			resultLength = int(palindromeLengthToRight[i]) - 1
		}
	}
	fmt.Println(palindromeLengthToRight)
	return s[resultStartIndex : resultStartIndex+resultLength]
}

func main() {
	fmt.Println(string(manacher([]byte("daaabaaa"))))
}
