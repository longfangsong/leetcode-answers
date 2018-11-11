package main

import (
	"fmt"
	"math"
	"strings"
)

func removeSame(board string, index int, putCount uint8) string {
	var left, right string
	if putCount == 1 {
		left = board[0:index]
		right = board[index+2:]
	} else if putCount == 2 {
		left = board[0:index]
		right = board[index+1:]
	}
	for (len(left) >= 2 && len(right) >= 1 && left[len(left)-1] == left[len(left)-2] && right[0] == left[len(left)-1]) ||
		(len(right) >= 2 && len(left) >= 1 && left[len(left)-1] == right[0] && right[0] == right[1]) {
		left = strings.TrimRight(left, right[0:1])
		right = strings.TrimLeft(right, right[0:1])
	}
	return left + right
}

func findMinStepImpl(board string, hand map[byte]uint8) int {
	if board == "" {
		return 0
	}
	if len(hand) == 0 {
		return -1
	}
	minStepCount := math.MaxInt64
	for index, ch := range board {
		if index < len(board)-1 {
			if board[index+1] == byte(ch) && hand[byte(ch)] >= 1 {
				// insert one ball
				hand[byte(ch)]--
				removed := removeSame(board, index, 1)
				result := findMinStepImpl(removed, hand)
				hand[byte(ch)]++
				if result != -1 && result+1 < minStepCount {
					minStepCount = result + 1
				}
			}
		}
		if hand[byte(ch)] >= 2 {
			// insert two balls
			hand[byte(ch)] -= 2
			removed := removeSame(board, index, 2)
			result := findMinStepImpl(removed, hand)
			hand[byte(ch)] += 2
			if result != -1 && result+2 < minStepCount {
				minStepCount = result + 2
			}
		}
	}
	if minStepCount == math.MaxInt64 {
		return -1
	}
	return minStepCount
}

func findMinStep(board string, hand string) int {
	handTypes := map[byte]uint8{}
	for _, ch := range hand {
		handTypes[byte(ch)]++
	}
	return findMinStepImpl(board, handTypes)
}

func main() {
	fmt.Println(findMinStep("W", "WW"))
}
