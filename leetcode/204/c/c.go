package main

// github.com/EndlessCheng/codeforces-go
func minDays(g [][]int) (ans int) {
	n, m := len(g), len(g[0])
	type pair struct{ x, y int }
	dir4 := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	vis := make([][]bool, n)
	for i := range vis {
		vis[i] = make([]bool, m)
	}
	var f func(i, j int)
	f = func(i, j int) {
		if i < 0 || i >= n || j < 0 || j >= m || vis[i][j] || g[i][j] == 0 {
			return
		}
		vis[i][j] = true
		for _, d := range dir4 {
			f(i+d.x, j+d.y)
		}
	}
	cc := func() (c int) {
		for i, row := range vis {
			for j := range row {
				vis[i][j] = false
			}
		}
		for i, row := range g {
			for j, v := range row {
				if !vis[i][j] && v == 1 {
					c++
					f(i, j)
				}
			}
		}
		return
	}
	if cc() != 1 {
		return
	}
	for i, gi := range g {
		for j, v := range gi {
			if v == 1 {
				g[i][j] = 0
				if cc() != 1 {
					return 1
				}
				g[i][j] = 1
			}
		}
	}
	return 2
}
