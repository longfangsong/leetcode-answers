package main

import (
	"fmt"
)

func minInt(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func maxInt(x int, y int) int {
	if x < y {
		return y
	}
	return x
}

func maxArea(height []int) int {
	frontCursor := 0
	backCursor := len(height) - 1
	maxArea := 0
	for frontCursor != backCursor {
		currentArea := minInt(height[frontCursor], height[backCursor]) * (backCursor - frontCursor)
		maxArea = maxInt(maxArea, currentArea)
		if height[frontCursor] < height[backCursor] {
			frontCursor++
		} else {
			backCursor--
		}
	}
	return maxArea
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
