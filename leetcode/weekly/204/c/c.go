package main

// github.com/EndlessCheng/codeforces-go
type pair struct{ x, y int }
var dir4 = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// O(nm) 求割点
func minDays(g [][]int) (ans int) {
	n, m := len(g), len(g[0])
	vis := make([][]bool, n)
	for i := range vis {
		vis[i] = make([]bool, m)
	}
	var dfs func(x, y int)
	dfs = func(x, y int) {
		vis[x][y] = true
		for _, d := range dir4 {
			if xx, yy := x+d.x, y+d.y; 0 <= xx && xx < n && 0 <= yy && yy < m && g[x][y] == 1 {
				if !vis[xx][yy] {
					dfs(xx, yy)
				}
			}
		}
	}
	found := false
	for i, row := range g {
		for j, v := range row {
			if v == 1 && !vis[i][j] {
				if found {
					return
				}
				found = true
				dfs(i, j)
			}
		}
	}
	if !found {
		return
	}

	dfn := make([][]int, n)
	for i := range dfn {
		dfn[i] = make([]int, m)
	}
	ts := 0
	var f func(x, y, fx, fy int) (int, bool)
	f = func(x, y, fx, fy int) (lowV int, cut bool) {
		ts++
		dfn[x][y] = ts
		lowV = ts
		childCnt := 0
		for _, d := range dir4 {
			if xx, yy := x+d.x, y+d.y; 0 <= xx && xx < n && 0 <= yy && yy < m && g[xx][yy] == 1 {
				if dfn[xx][yy] == 0 {
					childCnt++
					lowW, ct := f(xx, yy, x, y)
					if ct {
						return 0, true
					}
					if lowW >= dfn[x][y] {
						cut = true
					}
					lowV = min(lowV, lowW)
				} else if (xx != fx || yy != fy) && dfn[xx][yy] < dfn[x][y] {
					lowV = min(lowV, dfn[xx][yy])
				}
			}
		}
		if fx == -1 && fy == -1 && childCnt == 1 {
			cut = false
		}
		return
	}
o:
	for i, row := range g {
		for j, v := range row {
			if v == 1 {
				if _, cut := f(i, j, -1, -1); cut {
					return 1
				}
				break o
			}
		}
	}
	return 2
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
