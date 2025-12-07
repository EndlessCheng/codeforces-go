package main

import "sort"

// https://space.bilibili.com/206214
func totalScore(hp int, damage []int, requirement []int) (ans int64) {
	sum := make([]int, len(damage)+1)
	for i, req := range requirement {
		sum[i+1] = sum[i] + damage[i]
		low := sum[i+1] + req - hp
		j := sort.SearchInts(sum[:i+1], low)
		ans += int64(i - j + 1)
	}
	return
}
