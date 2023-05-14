package main

// https://space.bilibili.com/206214
func maxMoves(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没被计算过
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		if j == n-1 {
			return
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		for k := i - 1; k <= i+1; k++ {
			if 0 <= k && k < m && grid[k][j+1] > grid[i][j] {
				res = max(res, dfs(k, j+1)+1)
			}
		}
		*p = res // 记忆化
		return
	}
	for i := 0; i < m; i++ {
		ans = max(ans, dfs(i, 0))
	}
	return
}

func max(a, b int) int { if a < b { return b }; return a }
