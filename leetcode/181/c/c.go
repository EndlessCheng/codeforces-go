package main

func hasValidPath(grid [][]int) bool {
	n, m := len(grid), len(grid[0])

	// 右下左上
	dir4 := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	dir4ID := [6][2]int{{0, 2}, {1, 3}, {1, 2}, {0, 1}, {2, 3}, {0, 3}}
	g := make([][]int, n*m)
	for i, row := range grid {
		for j, street := range row {
			v := i*m + j
			for _, id := range dir4ID[street-1] {
				d := dir4[id]
				if x, y := i+d[0], j+d[1]; x >= 0 && x < n && y >= 0 && y < m {
					g[v] = append(g[v], x*m+y)
				} else {
					g[v] = append(g[v], -1)
				}
			}
		}
	}

	vis := make([]bool, n*m)
	var f func(v int) bool
	f = func(v int) bool {
		if v == n*m-1 {
			return true
		}
		vis[v] = true
		for _, w := range g[v] {
			if w != -1 && !vis[w] && (g[w][0] == v || g[w][1] == v) && f(w) {
				return true
			}
		}
		return false
	}
	return f(0)
}
