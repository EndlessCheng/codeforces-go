package main

import "strings"

// github.com/EndlessCheng/codeforces-go
func maximumSubsequenceCount(text, pattern string) (ans int64) {
	x, y := pattern[0], pattern[1]
	if x == y {
		c := int64(strings.Count(text, pattern[:1]))
		return c * (c + 1) / 2
	}
	cx, cy := 0, 0
	for i := range text {
		if ch := text[i]; ch == x {
			cx++
		} else if ch == y {
			cy++
			ans += int64(cx) // 每遇到一个 y，就累加左边 x 的个数
		}
	}
	return ans + int64(max(cx, cy))
}

func max(a, b int) int { if b > a { return b }; return a }
