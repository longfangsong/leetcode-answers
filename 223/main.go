package main

import (
	"sort"
)

type IntSlice []int

func (s IntSlice) Len() int           { return len(s) }
func (s IntSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}
func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	// 不相交的情况
	if C < E || D < F || G < A || H < B {
		return abs((C-A)*(B-D)) + abs((E-G)*(F-H))
	}
	// A,C,E,G取中间两个
	x := IntSlice{A, C, E, G}
	sort.Sort(x)
	width := x[2] - x[1]
	// B,D,F,H取中间两个
	y := IntSlice{B, D, F, H}
	sort.Sort(y)
	height := y[2] - y[1]
	return abs((C-A)*(B-D)) + abs((E-G)*(F-H)) - width*height
}

func main() {
	println(computeArea(-2, -2, 2, 2, 3, 3, 4, 4))
}
