package main

// github.com/EndlessCheng/codeforces-go
type pair struct{ x, y int }
var dir4 = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func containsCycle(g [][]byte) (ans bool) {
	n, m := len(g), len(g[0])
	vis := make([][]bool, n)
	for i := range vis {
		vis[i] = make([]bool, m)
	}
	var cur byte
	var f func(x, y, px, py int) bool
	f = func(x, y, px, py int) bool {
		vis[x][y] = true
		for _, d := range dir4 {
			if xx, yy := x+d.x, y+d.y; 0 <= xx && xx < n && 0 <= yy && yy < m && g[xx][yy] == cur {
				if vis[xx][yy] {
					if xx != px || yy != py {
						return true
					}
				} else if f(xx, yy, x, y) {
					return true
				}
			}
		}
		return false
	}
	for i, row := range g {
		for j := range row {
			if !vis[i][j] {
				cur = g[i][j]
				if f(i, j, -1, -1) {
					return true
				}
			}
		}
	}
	return
}
