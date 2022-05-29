package main

// https://space.bilibili.com/206214/dynamic
type pair struct{ x, y int }

var dir4 = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func minimumObstacles(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	dis := make([][]int, m)
	for i := range dis {
		dis[i] = make([]int, n)
		for j := range dis[i] {
			dis[i][j] = m * n
		}
	}
	dis[0][0] = 0
	q := [2][]pair{{{}}}
	for len(q[0]) > 0 || len(q[1]) > 0 {
		var p pair
		if len(q[0]) > 0 {
			p, q[0] = q[0][len(q[0])-1], q[0][:len(q[0])-1]
		} else {
			p, q[1] = q[1][0], q[1][1:]
		}
		for _, d := range dir4 {
			x, y := p.x+d.x, p.y+d.y
			if 0 <= x && x < m && 0 <= y && y < n {
				g := grid[x][y]
				if dis[p.x][p.y]+g < dis[x][y] {
					dis[x][y] = dis[p.x][p.y] + g
					q[g] = append(q[g], pair{x, y})
				}
			}
		}
	}
	return dis[m-1][n-1]
}
