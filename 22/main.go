package main

import "fmt"

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{""}
	}
	result := []string{}
	for parenthesisInCount := 0; parenthesisInCount < n; parenthesisInCount++ {
		for _, parenthesisIn := range generateParenthesis(parenthesisInCount) {
			for _, parenthesisOut := range generateParenthesis(n - 1 - parenthesisInCount) {
				result = append(result, "("+parenthesisIn+")"+parenthesisOut)
			}
		}
	}
	return result
}

func main() {
	fmt.Println(generateParenthesis(3))
}
