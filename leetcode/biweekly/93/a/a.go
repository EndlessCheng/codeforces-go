package main

import (
	"strconv"
)

// https://space.bilibili.com/206214
func maximumValue(strs []string) (ans int) {
	for _, s := range strs {
		x, err := strconv.Atoi(s)
		if err != nil {
			x = len(s)
		}
		ans = max(ans, x)
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
