package main

// https://space.bilibili.com/206214
var DIRS = [4][2]int{{1, 1}, {1, -1}, {-1, -1}, {-1, 1}}

func lenOfVDiagonal(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	memo := make([][][4][2]int, m)
	for i := range memo {
		memo[i] = make([][4][2]int, n)
	}

	var dfs func(int, int, int, int, int) int
	dfs = func(i, j, k, canTurn, target int) (res int) {
		i += DIRS[k][0]
		j += DIRS[k][1]
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != target {
			return
		}
		p := &memo[i][j][k][canTurn]
		if *p > 0 {
			return *p
		}
		defer func() { *p = res }()
		res = dfs(i, j, k, canTurn, 2-target) // 直行
		if canTurn == 1 {
			res = max(res, dfs(i, j, (k+1)%4, 0, 2-target)) // 右转
		}
		return res + 1
	}

	for i, row := range grid {
		for j, x := range row {
			if x == 1 {
				for k := range 4 { // 枚举起始方向
					ans = max(ans, dfs(i, j, k, 1, 2)+1)
				}
			}
		}
	}
	return
}
