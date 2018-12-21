package main

import "fmt"

var digitToCharacters = map[byte][]string{
	'2': []string{"a", "b", "c"},
	'3': []string{"d", "e", "f"},
	'4': []string{"g", "h", "i"},
	'5': []string{"j", "k", "l"},
	'6': []string{"m", "n", "o"},
	'7': []string{"p", "q", "r", "s"},
	'8': []string{"t", "u", "v"},
	'9': []string{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	if len(digits) == 1 {
		return digitToCharacters[digits[0]]
	}
	restletterCombinations := letterCombinations(digits[1:])
	result := []string{}
	for _, combination := range restletterCombinations {
		for _, firstCharacter := range digitToCharacters[digits[0]] {
			result = append(result, firstCharacter+combination)
		}
	}
	return result
}

func main() {
	fmt.Println(letterCombinations("23"))
}
