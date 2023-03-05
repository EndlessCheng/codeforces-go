package main

import "sort"

// https://space.bilibili.com/206214
func countWays(ranges [][]int) int {
	const mod int = 1e9 + 7
	sort.Slice(ranges, func(i, j int) bool { return ranges[i][0] < ranges[j][0] })
	ans, maxR := 2, ranges[0][1]
	for _, p := range ranges[1:] {
		if p[0] > maxR { // 产生了一个新的集合
			ans = ans * 2 % mod
		}
		maxR = max(maxR, p[1])
	}
	return ans
}

func max(a, b int) int { if a < b { return b }; return a }
