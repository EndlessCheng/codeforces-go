package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func maximumBags(capacity, rocks []int, additionalRocks int) (ans int) {
	for i := range capacity {
		capacity[i] -= rocks[i]
	}
	sort.Ints(capacity) // 先装剩余最小的
	for _, leftSpace := range capacity {
		if leftSpace > additionalRocks {
			break
		}
		ans++
		additionalRocks -= leftSpace
	}
	return
}
