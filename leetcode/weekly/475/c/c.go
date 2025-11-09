package main

import (
	"math"
	"slices"
)

// https://space.bilibili.com/206214
func maxPathScore1(grid [][]int, k int) int {
	n, m := len(grid[0]), len(grid)
	memo := make([][][]int, m)
	for i := range memo {
		memo[i] = make([][]int, n)
		for j := range memo[i] {
			memo[i][j] = make([]int, k+1)
			for p := range memo[i][j] {
				memo[i][j][p] = -1
			}
		}
	}

	var dfs func(int, int, int) int
	dfs = func(i, j, k int) int {
		if i < 0 || j < 0 || k < 0 { // 出界或者总花费超了
			return math.MinInt
		}
		if i == 0 && j == 0 {
			return 0 // 题目保证 grid[0][0] = 0
		}
		p := &memo[i][j][k]
		if *p != -1 {
			return *p
		}
		x := grid[i][j]
		if x > 0 {
			k--
		}
		res := max(dfs(i-1, j, k), dfs(i, j-1, k)) + x
		*p = res
		return res
	}

	ans := dfs(m-1, n-1, k)
	if ans < 0 {
		return -1
	}
	return ans
}

//

// 64. 最小路径和
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for j := 2; j <= n; j++ {
		f[0][j] = math.MaxInt
	}
	for i, row := range grid {
		f[i+1][0] = math.MaxInt
		for j, x := range row {
			f[i+1][j+1] = min(f[i+1][j], f[i][j+1]) + min(x, 1) // 值大于 0 的单元格花费 1
		}
	}
	return f[m][n]
}

func maxPathScore(grid [][]int, K int) int {
	if minPathSum(grid) > K {
		return -1
	}

	n, m := len(grid[0]), len(grid)
	K = min(K, m+n-2) // 至多花费 m+n-2
	f := make([][][]int, m+1)
	for i := range f {
		f[i] = make([][]int, n+1)
		for j := range f[i] {
			f[i][j] = make([]int, K+2)
			for p := range f[i][j] {
				f[i][j][p] = math.MinInt
			}
		}
	}
	f[0][1][1] = 0

	for i, row := range grid {
		for j, x := range row {
			for k := range min(K, i+j) + 1 { // 从 (0,0) 到 (i,j) 至多花费 i+j
				newK := k
				if x > 0 {
					newK--
				}
				f[i+1][j+1][k+1] = max(f[i][j+1][newK+1], f[i+1][j][newK+1]) + x
			}
		}
	}

	return slices.Max(f[m][n])
}
