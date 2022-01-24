package main

import "sort"

// 从大到小排序贪心

// github.com/EndlessCheng/codeforces-go
func minimumCost(cost []int) (ans int) {
	sort.Ints(cost)
	for i := len(cost) - 1; i >= 0; i -= 3 {
		ans += cost[i]
		if i > 0 {
			ans += cost[i-1]
		}
	}
	return
}
