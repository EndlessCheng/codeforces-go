package main

func domino(n int, m int, broken [][]int) (ans int) {
	valid := make([][]bool, n)
	for i := range valid {
		valid[i] = make([]bool, m)
		for j := range valid[i] {
			valid[i][j] = true
		}
	}
	for _, p := range broken {
		valid[p[0]][p[1]] = false
	}

	g := make([][]int, n*m)
	dir4 := [...][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for i, row := range valid {
		for j, ok := range row {
			if ok {
				v := i*m + j
				for _, d := range dir4 {
					if x, y := i+d[0], j+d[1]; x >= 0 && x < n && y >= 0 && y < m && valid[x][y] {
						w := x*m + y
						g[v] = append(g[v], w)
					}
				}
			}
		}
	}

	match := make([]int, n*m)
	for i := range match {
		match[i] = -1
	}
	var used []bool
	var f func(v int) bool
	f = func(v int) bool {
		used[v] = true
		for _, w := range g[v] {
			if mw := match[w]; mw == -1 || !used[mw] && f(mw) {
				match[w] = v
				match[v] = w
				return true
			}
		}
		return false
	}
	for v := range g {
		if match[v] == -1 {
			used = make([]bool, n*m)
			if f(v) {
				ans++
			}
		}
	}
	return
}
