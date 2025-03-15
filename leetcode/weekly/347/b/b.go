package main

import "math/bits"

// https://space.bilibili.com/206214
func differenceOfDistinctValues(grid [][]int) [][]int {
	m, n := len(grid), len(grid[0])
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}

	// 第一排在右上，最后一排在左下
	// 每排从左上到右下
	// 令 k=i-j+n，那么右上角 k=1，左下角 k=m+n-1
	for k := 1; k < m+n; k++ {
		// 核心：计算 j 的最小值和最大值
		minJ := max(n-k, 0)       // i=0 的时候，j=n-k，但不能是负数
		maxJ := min(m+n-1-k, n-1) // i=m-1 的时候，j=m+n-1-k，但不能超过 n-1

		set := uint(0)
		for j := minJ; j <= maxJ; j++ {
			i := k + j - n
			ans[i][j] = bits.OnesCount(set) // set 的大小
			set |= 1 << grid[i][j] // 把 grid[i][j] 加到 set 中
		}

		set = 0
		for j := maxJ; j >= minJ; j-- {
			i := k + j - n
			ans[i][j] = abs(ans[i][j] - bits.OnesCount(set))
			set |= 1 << grid[i][j]
		}
	}
	return ans
}

func abs(x int) int { if x < 0 { return -x }; return x }
