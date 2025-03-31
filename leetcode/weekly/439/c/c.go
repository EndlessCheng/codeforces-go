package main

import "math"

// https://space.bilibili.com/206214
func maxSum(nums []int, k, m int) int {
	n := len(nums)
	s := make([]int, n+1)
	for i, x := range nums {
		s[i+1] = s[i] + x // nums 的前缀和
	}

	f := make([]int, n+1)
	g := make([]int, n+1)
	for i := 1; i <= k; i++ {
		g[i*m-1] = math.MinInt
		mx := math.MinInt
		for j := i * m; j <= n-(k-i)*m; j++ {
			mx = max(mx, f[j-m]-s[j-m])
			g[j] = max(g[j-1], mx+s[j])
		}
		f, g = g, f
	}
	return f[n]
}
