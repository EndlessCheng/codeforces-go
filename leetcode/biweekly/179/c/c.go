package main

import "math"

// https://space.bilibili.com/206214
func minCost(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	// 异或和不会超过所有元素的 OR
	orAll := 0
	for _, row := range grid {
		for _, x := range row {
			orAll |= x
		}
	}

	vis := make([][][]bool, m)
	for i := range vis {
		vis[i] = make([][]bool, n)
		for j := range vis[i] {
			vis[i][j] = make([]bool, orAll+1)
		}
	}
	ans := math.MaxInt

	var dfs func(int, int, int)
	dfs = func(i, j, xor int) {
		// 最优性剪枝：如果答案已经最小（等于 0），那么不再搜索
		if ans == 0 || i < 0 || j < 0 || vis[i][j][xor] {
			return
		}
		vis[i][j][xor] = true
		xor ^= grid[i][j]
		if i == 0 && j == 0 {
			ans = min(ans, xor)
			return
		}
		dfs(i-1, j, xor)
		dfs(i, j-1, xor)
	}

	dfs(m-1, n-1, 0)
	return ans
}
