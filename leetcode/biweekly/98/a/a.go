package main

import (
	"strconv"
	"strings"
)

// https://space.bilibili.com/206214
func minMaxDifference(num int) int {
	mx, mn := 0, num
	s := strconv.Itoa(num)
	for _, c := range s {
		t := strings.ReplaceAll(s, string(c), "9")
		x, _ := strconv.Atoi(t)
		mx = max(mx, x)
		t = strings.ReplaceAll(s, string(c), "0")
		x, _ = strconv.Atoi(t)
		mn = min(mn, x)
	}
	return mx - mn
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
