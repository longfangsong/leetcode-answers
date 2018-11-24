package main

import (
	"bytes"
	"fmt"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	buffer := make([][]byte, numRows)
	nowInRowIndex := 0
	goDown := true
	for _, ch := range s {
		buffer[nowInRowIndex] = append(buffer[nowInRowIndex], byte(ch))
		if nowInRowIndex == numRows-1 {
			goDown = false
		} else if nowInRowIndex == 0 {
			goDown = true
		}
		if goDown {
			nowInRowIndex++
		} else {
			nowInRowIndex--
		}
	}
	return string(bytes.Join(buffer, []byte("")))
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
}
