package main

import "math/bits"

// https://space.bilibili.com/206214
func countKReducibleNumbers(s string, k int) (ans int) {
	const mod = 1_000_000_007
	n := len(s)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int, bool) int
	dfs = func(i, left1 int, isLimit bool) (res int) {
		if i == n {
			if !isLimit && left1 == 0 {
				return 1
			}
			return
		}
		if !isLimit {
			p := &memo[i][left1]
			if *p >= 0 {
				return *p
			}
			defer func() { *p = res }()
		}
		up := 1
		if isLimit {
			up = int(s[i] - '0')
		}
		for d := 0; d <= min(up, left1); d++ {
			res += dfs(i+1, left1-d, isLimit && d == up)
		}
		return res % mod
	}

	f := make([]int, n)
	for i := 1; i < n; i++ {
		f[i] = f[bits.OnesCount(uint(i))] + 1
		if f[i] <= k {
			// 计算有多少个二进制数恰好有 i 个 1
			ans += dfs(0, i, true)
		}
	}
	return ans % mod
}
