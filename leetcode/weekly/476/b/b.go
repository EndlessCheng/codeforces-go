package main

import "strings"

// https://space.bilibili.com/206214
func minLengthAfterRemovals(s string) int {
	k := strings.Count(s, "a")
	return abs(k*2 - len(s))
}

func abs(x int) int { if x < 0 { return -x }; return x }
