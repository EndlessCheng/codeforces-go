package main

import "math"

// https://space.bilibili.com/206214
func findSafeWalk(grid [][]int, health int) bool {
	type pair struct{ x, y int }
	dirs := []pair{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	m, n := len(grid), len(grid[0])
	dis := make([][]int, m)
	for i := range dis {
		dis[i] = make([]int, n)
		for j := range dis[i] {
			dis[i][j] = math.MaxInt
		}
	}
	dis[0][0] = grid[0][0]
	q := [2][]pair{{{}}} // 两个 slice 头对头来实现 deque
	for {
		var p pair
		if len(q[0]) > 0 {
			p, q[0] = q[0][len(q[0])-1], q[0][:len(q[0])-1]
		} else {
			p, q[1] = q[1][0], q[1][1:]
		}
		i, j := p.x, p.y
		if dis[i][j] >= health {
			return false
		}
		if i == m-1 && j == n-1 {
			return true
		}
		for _, d := range dirs {
			x, y := i+d.x, j+d.y
			if 0 <= x && x < m && 0 <= y && y < n {
				cost := grid[x][y]
				if dis[i][j]+cost < dis[x][y] {
					dis[x][y] = dis[i][j] + cost
					q[cost] = append(q[cost], pair{x, y})
				}
			}
		}
	}
}

func findSafeWalk2(grid [][]int, health int) bool {
	type pair struct{ x, y int }
	dirs := []pair{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	m, n := len(grid), len(grid[0])
	dis := make([][]int, m)
	for i := range dis {
		dis[i] = make([]int, n)
		for j := range dis[i] {
			dis[i][j] = math.MaxInt
		}
	}
	dis[0][0] = grid[0][0]
	q := [2][]pair{{{}}} // 两个 slice 头对头来实现 deque
	for len(q[0]) > 0 || len(q[1]) > 0 {
		var p pair
		if len(q[0]) > 0 {
			p, q[0] = q[0][len(q[0])-1], q[0][:len(q[0])-1]
		} else {
			p, q[1] = q[1][0], q[1][1:]
		}
		i, j := p.x, p.y
		for _, d := range dirs {
			x, y := i+d.x, j+d.y
			if 0 <= x && x < m && 0 <= y && y < n {
				cost := grid[x][y]
				if dis[i][j]+cost < dis[x][y] {
					dis[x][y] = dis[i][j] + cost
					q[cost] = append(q[cost], pair{x, y})
				}
			}
		}
	}
	return dis[m-1][n-1] < health
}
