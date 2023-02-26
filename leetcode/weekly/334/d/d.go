package main

import (
	"sort"
)

// https://space.bilibili.com/206214
type pair struct{ x, y int }

var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func minimumTime(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	if grid[0][1] > 1 && grid[1][0] > 1 { // 无法「等待」
		return -1
	}

	vis := make([][]int, m)
	for i := range vis {
		vis[i] = make([]int, n)
	}
	endTime := sort.Search(1e5+m+n, func(endTime int) bool {
		if endTime < grid[m-1][n-1] || endTime < m+n-2 {
			return false
		}
		vis[m-1][n-1] = endTime
		q := []pair{{m - 1, n - 1}}
		for t := endTime - 1; len(q) > 0; t-- {
			tmp := q
			q = nil
			for _, p := range tmp {
				for _, d := range dirs { // 枚举周围四个格子
					x, y := p.x+d.x, p.y+d.y
					if 0 <= x && x < m && 0 <= y && y < n && vis[x][y] != endTime && grid[x][y] <= t {
						if x == 0 && y == 0 {
							return true
						}
						vis[x][y] = endTime // 用二分的值来标记，避免重复创建 vis 数组
						q = append(q, pair{x, y})
					}
				}
			}
		}
		return false
	})
	return endTime + (endTime+m+n)%2
}

func min(a, b int) int { if a > b { return b }; return a }
