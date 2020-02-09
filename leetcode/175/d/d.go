package main

func maxStudents(mat [][]byte) (ans int) {
	dir4 := [4][2]int{{0, -1}, {0, 1}, {-1, -1}, {-1, 1}}
	n, m := len(mat), len(mat[0])
	validSeatCnt := 0
	g := make([][]int, n*m)
	for i, mi := range mat {
		for j, mij := range mi {
			if mij == '.' {
				validSeatCnt++
				v := i*m + j
				for _, d := range dir4 {
					if x, y := i+d[0], j+d[1]; x >= 0 && x < n && y >= 0 && y < m && mat[x][y] == '.' {
						g[v] = append(g[v], x*m+y)
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
	ans = validSeatCnt - ans
	return
}
