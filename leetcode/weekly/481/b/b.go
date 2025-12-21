package main

import "slices"

// https://space.bilibili.com/206214
func minCost(s string, cost []int) int64 {
	total := 0
	sum := [26]int{}
	for i, x := range cost {
		total += x
		sum[s[i]-'a'] += x
	}
	return int64(total - slices.Max(sum[:]))
}
