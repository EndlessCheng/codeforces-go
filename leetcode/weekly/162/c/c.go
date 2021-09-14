package main

func closedIsland(grid [][]int) (aa int) {
	v := [200][200]bool{}
	var dfs func(int, int) bool
	n := len(grid)
	m := len(grid[0])
	var d4 = [...][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	dfs = func(i, j int) bool {
		if i < 0 || i >= n || j < 0 || j >= m {
			return false
		}
		if grid[i][j] == 1 {
			return true
		}
		if v[i][j] {
			return true
		}
		v[i][j] = true
		// FIXME: 反思：太久没写简单的 DFS 了导致没遍历完连通分量就提前 return 了
		res := true
		for _, dir := range d4 {
			if !dfs(i+dir[0], j+dir[1]) {
				res = false
			}
		}
		return res
	}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 && !v[i][j] {
				if dfs(i, j) {
					aa++
				}
			}
		}
	}
	return
}

