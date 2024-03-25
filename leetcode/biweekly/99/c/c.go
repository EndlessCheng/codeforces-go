package main

import "slices"

// https://space.bilibili.com/206214
func countWays(ranges [][]int) int {
	slices.SortFunc(ranges, func(p, q []int) int { return p[0] - q[0] })
	ans, maxR := 1, -1
	for _, p := range ranges {
		if p[0] > maxR { // 无法合并
			ans = ans * 2 % 1_000_000_007 // 新区间
		}
		maxR = max(maxR, p[1]) // 合并
	}
	return ans
}
