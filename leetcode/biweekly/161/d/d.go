package main

import (
	"math/bits"
	"strconv"
)

// https://space.bilibili.com/206214
func popcountDepth(n int64, k int) (ans int64) {
	if k == 0 {
		return 1
	}

	// 注：也可以不转成字符串，下面 dfs 用位运算取出 n 的第 i 位
	// 但转成字符串的通用性更好
	s := strconv.FormatInt(n, 2)
	m := len(s)
	if k == 1 {
		return int64(m - 1)
	}

	memo := make([][]int64, m)
	for i := range memo {
		memo[i] = make([]int64, m+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(int, int, bool) int64
	dfs = func(i, left1 int, isLimit bool) (res int64) {
		if i == m {
			if left1 == 0 {
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
		return
	}

	f := make([]int, m+1)
	for i := 1; i <= m; i++ {
		f[i] = f[bits.OnesCount(uint(i))] + 1
		if f[i] == k {
			// 计算有多少个二进制数恰好有 i 个 1
			ans += dfs(0, i, true)
		}
	}
	return
}
