package main

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	begin := 0
	end := 1
	currentMaxLength := 1
	alreadyHaveCharacter := map[byte]bool{}
	alreadyHaveCharacter[s[begin]] = true
	// 选区为s[begin,end)，这一区间内无重复字符
	for end != len(s) {
		for alreadyHaveCharacter[s[end]] && begin != end {
			alreadyHaveCharacter[s[begin]] = false
			begin++
		}
		alreadyHaveCharacter[s[end]] = true
		end++
		if end-begin > currentMaxLength {
			currentMaxLength = end - begin
		}
	}
	return currentMaxLength
}
func main() {
	println(lengthOfLongestSubstring("a"))
}
