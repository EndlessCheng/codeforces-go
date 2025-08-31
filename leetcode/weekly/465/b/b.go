package main

import (
	"math"
	"slices"
)

// https://space.bilibili.com/206214
const mx = 100_001

var divisors [mx][]int

func init() {
	// 预处理每个数的因子
	for i := 1; i < mx; i++ {
		for j := i; j < mx; j += i {
			divisors[j] = append(divisors[j], i)
		}
	}
}

func minDifference(n, k int) (ans []int) {
	minDiff := math.MaxInt
	path := make([]int, k)
	var dfs func(int, int, int, int)
	dfs = func(i, n, mn, mx int) {
		if i == k-1 {
			d := max(mx, n) - min(mn, n) // 最后一个数是 n
			if d < minDiff {
				minDiff = d
				path[i] = n
				ans = slices.Clone(path)
			}
			return
		}
		for _, d := range divisors[n] { // 枚举 x 的因子
			path[i] = d // 直接覆盖，无需恢复现场
			dfs(i+1, n/d, min(mn, d), max(mx, d))
		}
	}
	dfs(0, n, math.MaxInt, 0)
	return
}
