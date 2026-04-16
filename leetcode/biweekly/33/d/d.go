package main

// github.com/EndlessCheng/codeforces-go
var dirs = []struct{ x, y int }{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} // 左右上下

func containsCycle(grid [][]byte) bool {
	m, n := len(grid), len(grid[0])
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}

	var dfs func(int, int, int, int) bool
	dfs = func(x, y, px, py int) bool {
		vis[x][y] = true
		for _, d := range dirs { // 枚举移动方向
			i, j := x+d.x, y+d.y
			if (i != px || j != py) && // (i, j) 不是上一步的格子 (px, py)
				0 <= i && i < m && 0 <= j && j < n && // (i, j) 没有出界
				grid[i][j] == grid[x][y] && // (i, j) 和 (x, y) 的格子值相同
				(vis[i][j] || dfs(i, j, x, y)) { // 如果之前访问过 (i, j)，那么找到了环，否则继续递归找
				return true
			}
		}
		return false
	}

	for i, row := range vis {
		for j, b := range row {
			if !b && dfs(i, j, -1, -1) {
				return true
			}
		}
	}
	return false
}
