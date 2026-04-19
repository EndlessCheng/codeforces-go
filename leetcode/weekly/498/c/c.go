package main

import "slices"

// https://space.bilibili.com/206214
var dirs = []struct{ x, y int }{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} // 左右上下

func colorGrid(n, m int, sources [][]int) [][]int {
	slices.SortFunc(sources, func(a, b []int) int { return b[2] - a[2] })

	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, m)
	}
	for _, p := range sources {
		ans[p[0]][p[1]] = p[2] // 初始颜色
	}

	q := sources
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		x, y, c := p[0], p[1], p[2]
		for _, d := range dirs { // 向四个方向扩散
			i, j := x+d.x, y+d.y
			if 0 <= i && i < n && 0 <= j && j < m && ans[i][j] == 0 { // (i, j) 未着色
				ans[i][j] = c // 着色
				q = append(q, []int{i, j, c}) // 继续扩散
			}
		}
	}

	return ans
}
