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
		for j := i; j < mx; j += i { // 枚举 i 的倍数 j
			divisors[j] = append(divisors[j], i) // i 是 j 的因子
		}
	}
}

func minDifference(n, k int) (ans []int) {
	minDiff := math.MaxInt
	path := make([]int, k)
	var dfs func(int, int)
	dfs = func(i, n int) {
		if i == k-1 {
			// path[0] 最小，n 最大
			if n-path[0] < minDiff {
				minDiff = n - path[0]
				path[i] = n
				ans = slices.Clone(path)
			}
			return
		}
		for _, d := range divisors[n] { // 枚举 x 的因子
			if d*d > n || i > 0 && d-path[0] >= minDiff {
				break
			}
			if i == 0 || d >= path[i-1] {
				path[i] = d // 直接覆盖，无需恢复现场
				dfs(i+1, n/d)
			}
		}
	}
	dfs(0, n)
	return
}
