package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func maximumMinimumPath(a [][]int) (ans int) {
	type pair struct{ x, y int }
	dir4 := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	n, m := len(a), len(a[0])
	ans = sort.Search(min(a[0][0], a[n-1][m-1])+1, func(low int) bool {
		vis := make([][]bool, n)
		for i := range vis {
			vis[i] = make([]bool, m)
		}
		var f func(int, int) bool
		f = func(x, y int) bool {
			if x < 0 || x >= n || y < 0 || y >= m || vis[x][y] || a[x][y] < low {
				return false
			}
			if x == n-1 && y == m-1 {
				return true
			}
			vis[x][y] = true
			for _, d := range dir4 {
				if f(x+d.x, y+d.y) {
					return true
				}
			}
			return false
		}
		return !f(0, 0)
	}) - 1
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
