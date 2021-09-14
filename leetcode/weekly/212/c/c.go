package main

import "sort"

// 另外两个方法：
// - 从小到大加入并查集直至起终点在同一个集合
// - dij，把 + 换成 max

// 从码量上来说最小的还是二分

// github.com/EndlessCheng/codeforces-go
func minimumEffortPath(a [][]int) (ans int) {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	type pair struct{ x, y int }
	dir4 := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	n, m := len(a), len(a[0])
	return sort.Search(1e6, func(maxD int) bool {
		vis := make([][]bool, n)
		for i := range vis {
			vis[i] = make([]bool, m)
		}
		vis[0][0] = true
		q := []pair{{}}
		for len(q) > 0 {
			p := q[0]
			q = q[1:]
			x, y := p.x, p.y
			if x == n-1 && y == m-1 {
				return true
			}
			for _, d := range dir4 {
				if xx, yy := x+d.x, y+d.y; 0 <= xx && xx < n && 0 <= yy && yy < m && !vis[xx][yy] && abs(a[xx][yy]-a[x][y]) <= maxD {
					vis[xx][yy] = true
					q = append(q, pair{xx, yy})
				}
			}
		}
		return false
	})
}
