package main

import "slices"

// github.com/EndlessCheng/codeforces-go
func minimumCost(cost []int) (ans int) {
	slices.Sort(cost)
	// 从大到小，买两个送一个
	for i := len(cost) - 1; i >= 0; i -= 3 {
		ans += cost[i]
		if i > 0 {
			ans += cost[i-1]
		}
	}
	return
}
