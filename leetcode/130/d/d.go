package main

func numEnclaves(g [][]int) (ans int) {
	n, m := len(g), len(g[0])
	vis := make([][]bool, n)
	for i := range vis {
		vis[i] = make([]bool, m)
	}
	type pair struct{ x, y int }
	dir4 := [...]pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右
	var f func(i, j int)
	f = func(i, j int) {
		if vis[i][j] || g[i][j] == 0 {
			return
		}
		vis[i][j] = true
		for _, d := range dir4 {
			if xx, yy := i+d.x, j+d.y; 0 <= xx && xx < n && 0 <= yy && yy < m {
				f(xx, yy)
			}
		}
	}
	for i, gi := range g {
		for j := range gi {
			if i == 0 || i == n-1 || j == 0 || j == m-1 {
				f(i, j)
			}
		}
	}
	for i, row := range g {
		for j, v := range row {
			if v == 1 && !vis[i][j] {
				ans++
			}
		}
	}
	return
}
