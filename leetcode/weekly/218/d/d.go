package main

import "math/bits"

// github.com/EndlessCheng/codeforces-go
func minimumIncompatibility(a []int, k int) int {
	n := len(a)
	dp := make([]int, 1<<n)
	cache := make([]int, 1<<n)
	for i := range dp {
		dp[i] = -1
		cache[i] = -1
	}
	const inf int = 1e9
	calc := func(sub uint) int {
		res := cache[sub]
		if res != -1 {
			return res
		}
		mi, mx := inf, 0
		vis := make([]bool, n+1)
		for ; sub > 0; sub &= sub - 1 {
			v := a[bits.TrailingZeros(sub)]
			if vis[v] {
				res = inf
				return res
			}
			mi = min(mi, v)
			mx = max(mx, v)
			vis[v] = true
		}
		res = mx - mi
		return res
	}

	k = n / k
	var f func(int) int
	f = func(set int) (res int) {
		if set == 0 {
			return
		}
		dv := &dp[set]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		res = inf
		sub := set
		for ok := true; ok; ok = sub != set {
			if bits.OnesCount(uint(sub)) == k {
				if d := calc(uint(sub)); d < inf {
					res = min(res, d+f(set^sub))
				}
			}
			sub = (sub - 1) & set
		}
		return
	}
	ans := f(1<<n - 1)
	if ans >= inf {
		ans = -1
	}
	return ans
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
