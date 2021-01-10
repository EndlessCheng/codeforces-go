package main

import (
	"math/bits"
	"sort"
)

// 0 ms (https://leetcode.com/problems/find-minimum-time-to-finish-all-jobs/discuss/1010057/Python-Binary-search-24ms)
func minimumTimeRequired(a []int, k int) int {
	sort.Ints(a)
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sort.Search(sum, func(mx int) bool {
		cap := make([]int, k)
		for i := range cap {
			cap[i] = mx
		}
		var f func(int) bool
		f = func(p int) bool {
			if p < 0 {
				return true
			}
			for i := k - 1; i >= 0; i-- {
				if cap[i] >= a[p] {
					cap[i] -= a[p]
					if f(p - 1) {
						return true
					}
					cap[i] += a[p]
				}
				if cap[i] == mx { // 每人至少要分配一个任务
					break
				}
			}
			return false
		}
		return f(len(a) - 1)
	})
}

func minimumTimeRequiredDP(a []int, k int) int {
	m := 1 << len(a)
	sum := make([]int, m)
	for i := 1; i < m; i++ {
		p := bits.TrailingZeros(uint(i))
		sum[i] = sum[i&^(1<<p)] + a[p]
	}
	dp := make([][]int, k)
	for i := range dp {
		dp[i] = make([]int, m)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var f func(int, int) int
	f = func(p, set int) (res int) {
		if set == 0 {
			return
		}
		if p == k {
			return 1e9
		}
		dv := &dp[p][set]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		res = 1e9
		for sub, ok := set, true; ok; ok = sub != set {
			res = min(res, max(sum[sub], f(p+1, set^sub)))
			sub = (sub - 1) & set
		}
		return
	}
	return f(0, m-1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
