package main

/*
一开始想复杂了，给 $\textit{grid1}$ 的每个岛屿都求了个连通分量，写了一长串代码。

其实注意到只要 $\textit{grid2}$ 的岛屿在 $\textit{grid1}$ 对应位置上不存在水域，那其就是子岛屿。
*/

// github.com/EndlessCheng/codeforces-go
var dir4 = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
var ok bool

func countSubIslands(grid1, grid2 [][]int) (ans int) {
	n, m := len(grid2), len(grid2[0])
	var dfs func(int, int)
	dfs = func(x, y int) {
		grid2[x][y] = 0
		if grid1[x][y] == 0 { // 找到水域
			ok = false
		}
		for _, d := range dir4 {
			if x, y := x+d.x, y+d.y; 0 <= x && x < n && 0 <= y && y < m && grid2[x][y] > 0 {
				dfs(x, y)
			}
		}
	}
	for i, row := range grid2 {
		for j, v := range row {
			if v > 0 {
				ok = true
				dfs(i, j)
				if ok {
					ans++
				}
			}
		}
	}
	return
}
