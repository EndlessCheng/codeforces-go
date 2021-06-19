package main

// github.com/EndlessCheng/codeforces-go
type pair struct{ x, y int }
var dir4 = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func largestArea(grid []string) (ans int) {
	n, m := len(grid), len(grid[0])
	vis := make([][]bool, n)
	for i := range vis {
		vis[i] = make([]bool, m)
	}
	var cnt int
	var cur byte
	var f func(int, int) bool
	f = func(x, y int) bool {
		if x < 0 || x >= n || y < 0 || y >= m || grid[x][y] == '0' {
			return false
		}
		if vis[x][y] || grid[x][y] != cur {
			return true
		}
		cnt++
		vis[x][y] = true
		ok := true
		for _, d := range dir4 {
			if !f(x+d.x, y+d.y) {
				ok = false
			}
		}
		return ok
	}
	for i, row := range grid {
		for j, v := range row {
			if v != '0' && !vis[i][j] {
				cnt, cur = 0, byte(v)
				if f(i, j) && cnt > ans {
					ans = cnt
				}
			}
		}
	}
	return
}
