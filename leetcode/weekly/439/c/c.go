package main

import "math"

// https://space.bilibili.com/206214
func maxSum(nums []int, k, m int) int {
	n := len(nums)
	s := make([]int, n+1)
	for i, x := range nums {
		s[i+1] = s[i] + x
	}

	f := make([]int, n+1)
	d := make([]int, n+1)
	for i := 1; i <= k; i++ {
		for j := i*m - m; j < i*m; j++ {
			d[j] = f[j] - s[j]
			f[j] = math.MinInt / 2 // 即使 [0,j) 全选，也没有 i 个长为 m 的子数组
		}
		mx := math.MinInt
		// 左右两边留出足够空间给其他子数组
		for j := i * m; j <= n-(k-i)*m; j++ {
			// mx 表示最大的 f[L]-s[L]，其中 L 在区间 [(i-1)*m, j-m] 中
			mx = max(mx, d[j-m])
			d[j] = f[j] - s[j]
			f[j] = max(f[j-1], mx+s[j]) // 不选 vs 选
		}
	}
	return f[n]
}

func maxSum2(nums []int, k, m int) int {
	n := len(nums)
	s := make([]int, n+1)
	for i, x := range nums {
		s[i+1] = s[i] + x
	}

	f := make([]int, n+1)
	for i := 1; i <= k; i++ {
		nf := make([]int, n+1)
		for j := range nf {
			nf[j] = math.MinInt / 2
		}
		mx := math.MinInt
		// 左右两边留出足够空间给其他子数组
		for j := i * m; j <= n-(k-i)*m; j++ {
			// mx 表示最大的 f[L]-s[L]，其中 L 在区间 [(i-1)*m, j-m] 中
			mx = max(mx, f[j-m]-s[j-m])
			nf[j] = max(nf[j-1], mx+s[j]) // 不选 vs 选
		}
		f = nf
	}
	return f[n]
}

func maxSum1(nums []int, k, m int) int {
	n := len(nums)
	s := make([]int, n+1)
	for i, x := range nums {
		s[i+1] = s[i] + x
	}

	f := make([][]int, k+1)
	f[0] = make([]int, n+1)
	for i := 1; i <= k; i++ {
		f[i] = make([]int, n+1)
		for j := range f[i] {
			f[i][j] = math.MinInt / 2
		}
		mx := math.MinInt
		// 左右两边留出足够空间给其他子数组
		for j := i * m; j <= n-(k-i)*m; j++ {
			// mx 表示最大的 f[i-1][L]-s[L]，其中 L 在区间 [(i-1)*m, j-m] 中
			mx = max(mx, f[i-1][j-m]-s[j-m])
			f[i][j] = max(f[i][j-1], mx+s[j]) // 不选 vs 选
		}
	}
	return f[k][n]
}
