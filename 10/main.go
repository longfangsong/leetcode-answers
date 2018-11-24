package main

import (
	"fmt"
)

func isMatchImpl(s string, p string, knownUnmatch map[string]bool) bool {
	if knownUnmatch[s+"|"+p] {
		return false
	}
	if s == p {
		return true
	} else if s == "" {
		// 形如a*b*c*合法，其他情况都不合法
		if len(p)%2 == 1 {
			// 奇数长一定不合法
			return false
		}
		for i := 1; i < len(p); i += 2 {
			if p[i] != '*' {
				knownUnmatch[s+"|"+p] = true
				return false
			}
		}
		return true
	} else if p == "" {
		return false
	}
	if len(p) >= 2 && p[1] == '*' {
		// p[0]* 啥都不匹配
		if isMatchImpl(s, p[2:], knownUnmatch) {
			return true
		}
		// p[0]* 匹配了一个以上的字符
		if (s[0] == p[0] || p[0] == '.') && isMatchImpl(s[1:], p, knownUnmatch) {
			return true
		}
		knownUnmatch[s+"|"+p] = true
		return false
	}
	if (s[0] == p[0] || p[0] == '.') && isMatchImpl(s[1:], p[1:], knownUnmatch) {
		return true
	} else {
		knownUnmatch[s+"|"+p] = true
		return false
	}
}

func isMatch(s string, p string) bool {
	return isMatchImpl(s, p, map[string]bool{})
}

func main() {
	fmt.Println(isMatch("aaaaaaab", "a*a*a*a*c"))
}
