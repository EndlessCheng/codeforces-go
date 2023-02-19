package main

import "sort"

// https://space.bilibili.com/206214
func minimizeSum(a []int) int {
	sort.Ints(a)
	n := len(a)
	return min(min(a[n-3]-a[0], a[n-2]-a[1]), a[n-1]-a[2])
}

func min(a, b int) int { if a > b { return b }; return a }

