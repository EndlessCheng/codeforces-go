package main

import "strings"

// https://space.bilibili.com/206214
func minimumRecolors(s string, k int) int {
	n := len(s)
	cntW := strings.Count(s[:k], "W")
	ans := cntW
	for i := k; i < n; i++ {
		cntW += int(s[i]&1) - int(s[i-k]&1)
		ans = min(ans, cntW)
	}
	return ans
}

func min(a, b int) int { if a > b { return b }; return a }
