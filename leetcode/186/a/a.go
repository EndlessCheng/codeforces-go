package main

import "strings"

func maxScore(s string) (ans int) {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 1; i < len(s); i++ {
		ans = max(ans, strings.Count(s[:i], "0")+strings.Count(s[i:], "1"))
	}
	return
}
