package main

var dirs = [7][2][2]int{
	{},
	{{0, -1}, {0, 1}},  // 站在街道 1，可以往左或者往右
	{{-1, 0}, {1, 0}},  // 站在街道 2，可以往上或者往下
	{{0, -1}, {1, 0}},  // 站在街道 3，可以往左或者往下
	{{0, 1}, {1, 0}},   // 站在街道 4，可以往右或者往下
	{{0, -1}, {-1, 0}}, // 站在街道 5，可以往左或者往上
	{{0, 1}, {-1, 0}},  // 站在街道 6，可以往右或者往上
}

// 判断街道 street 是否包含移动方向 dir
func contains(street int, dir [2]int) bool {
	// 也可以写 slices.Contains(dirs[street][:], dir)
	return dirs[street][0] == dir || dirs[street][1] == dir
}

func hasValidPath(grid [][]int) bool {
	m, n := len(grid), len(grid[0])
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}

	var dfs func(int, int) bool
	dfs = func(x, y int) bool {
		if x == m-1 && y == n-1 {
			return true
		}
		vis[x][y] = true // 标记 (x, y) 访问过，从而避免重复访问
		for _, d := range dirs[grid[x][y]] { // 枚举下一步往哪走
			i, j := x+d[0], y+d[1]
			if 0 <= i && i < m && 0 <= j && j < n && !vis[i][j] &&
				contains(grid[i][j], [2]int{-d[0], -d[1]}) && dfs(i, j) {
				return true
			}
		}
		return false
	}

	return dfs(0, 0)
}
