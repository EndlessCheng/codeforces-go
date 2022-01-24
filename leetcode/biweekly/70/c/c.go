package main

import "sort"

// github.com/EndlessCheng/codeforces-go
var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func highestRankedKItems(grid [][]int, pricing, start []int, k int) (ans [][]int) {
	m, n := len(grid), len(grid[0])
	low, high := pricing[0], pricing[1]
	sx, sy := start[0], start[1]
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	vis[sx][sy] = true
	q := [][]int{{sx, sy}}
	for len(q) > 0 { // 分层 BFS
		// 此时 q 内所有元素到起点的距离均相同，因此按照题目中的第 2~4 关键字排序后，就可以将价格在 [low,high] 内的位置加入答案
		sort.Slice(q, func(i, j int) bool {
			ax, ay, bx, by := q[i][0], q[i][1], q[j][0], q[j][1]
			pa, pb := grid[ax][ay], grid[bx][by]
			return pa < pb || pa == pb && (ax < bx || ax == bx && ay < by)
		})
		l := sort.Search(len(q), func(i int) bool { return grid[q[i][0]][q[i][1]] >= low })
		r := sort.Search(len(q), func(i int) bool { return grid[q[i][0]][q[i][1]] > high })
		ans = append(ans, q[l:r]...)
		if len(ans) >= k {
			return ans[:k]
		}
		tmp := q
		q = nil
		for _, p := range tmp {
			for _, d := range dirs {
				if x, y := p[0]+d.x, p[1]+d.y; 0 <= x && x < m && 0 <= y && y < n && !vis[x][y] && grid[x][y] != 0 {
					vis[x][y] = true
					q = append(q, []int{x, y})
				}
			}
		}
	}
	return
}
