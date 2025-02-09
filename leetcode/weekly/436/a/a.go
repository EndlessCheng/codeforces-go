package main

import "slices"

// https://space.bilibili.com/206214
func sortMatrix(grid [][]int) [][]int {
	m, n := len(grid), len(grid[0])
	// 第一排在右上，最后一排在左下
	// 每排从左上到右下
	// 令 k = i-j+n，那么右上角 k=1，左下角 k=m+n-1
	for k := 1; k < m+n; k++ {
		// 核心：计算 j 的最小值和最大值
		minJ := max(n-k, 0)       // i=0 的时候，j=n-k，但不能是负数
		maxJ := min(m+n-1-k, n-1) // i=m-1 的时候，j=m+n-1-k，但不能超过 n-1
		a := []int{}
		for j := minJ; j <= maxJ; j++ {
			a = append(a, grid[k+j-n][j]) // 根据 k 的定义计算 i
		}
		if minJ > 0 {
			slices.Sort(a)
		} else {
			slices.SortFunc(a, func(a, b int) int { return b - a })
		}
		for j := minJ; j <= maxJ; j++ {
			grid[k+j-n][j] = a[j-minJ]
		}
	}
	return grid
}
