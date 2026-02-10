package main

import "math"

// https://space.bilibili.com/206214
func maxScore1(nums1, nums2 []int, K int) int64 {
	n, m := len(nums1), len(nums2)
	f := make([][][]int, K+1)
	for k := range f {
		f[k] = make([][]int, n+1)
		for i := range f[k] {
			f[k][i] = make([]int, m+1)
			if k > 0 {
				for j := range f[k][i] {
					f[k][i][j] = math.MinInt
				}
			}
		}
	}
	for k := 1; k <= K; k++ {
		for i := k - 1; i < n-(K-k); i++ { // 后面还要选 K-k 个下标对
			for j := k - 1; j < m-(K-k); j++ {
				f[k][i+1][j+1] = max(f[k][i][j+1], f[k][i+1][j], f[k-1][i][j]+nums1[i]*nums2[j])
			}
		}
	}
	return int64(f[K][n][m])
}

func maxScore(nums1, nums2 []int, K int) int64 {
	n, m := len(nums1), len(nums2)
	f := make([][]int, n+1)
	g := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, m+1)
		g[i] = make([]int, m+1)
	}
	for k := 1; k <= K; k++ {
		for _, row := range g {
			for j := range row {
				row[j] = math.MinInt
			}
		}
		for i := k - 1; i < n-(K-k); i++ { // 后面还要选 K-k 个下标对
			for j := k - 1; j < m-(K-k); j++ {
				g[i+1][j+1] = max(g[i][j+1], g[i+1][j], f[i][j]+nums1[i]*nums2[j])
			}
		}
		f, g = g, f
	}
	return int64(f[n][m])
}
