package main

import (
	"math"
	"unicode"
)

// https://space.bilibili.com/206214
func minMoves(matrix []string) int {
	m, n := len(matrix), len(matrix[0])
	if matrix[m-1][n-1] == '#' {
		return -1
	}

	type pair struct{ x, y int }
	pos := ['Z' + 1][]pair{}
	for i, row := range matrix {
		for j, c := range row {
			if unicode.IsUpper(c) {
				pos[c] = append(pos[c], pair{i, j})
			}
		}
	}

	dirs := []pair{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	dis := make([][]int, m)
	for i := range dis {
		dis[i] = make([]int, n)
		for j := range dis[i] {
			dis[i][j] = math.MaxInt
		}
	}
	dis[0][0] = 0

	// 两个 slice 头对头，模拟 deque
	q0 := []pair{{}}
	q1 := []pair{}

	for len(q0) > 0 || len(q1) > 0 {
		// 弹出队首
		var p pair
		if len(q0) > 0 {
			p, q0 = q0[len(q0)-1], q0[:len(q0)-1]
		} else {
			p, q1 = q1[0], q1[1:]
		}

		if p.x == m-1 && p.y == n-1 {
			return dis[p.x][p.y]
		}

		if c := matrix[p.x][p.y]; c != '.' {
			// 使用所有传送门
			for _, q := range pos[c] {
				x, y := q.x, q.y
				if dis[p.x][p.y] < dis[x][y] {
					dis[x][y] = dis[p.x][p.y]
					q0 = append(q0, pair{x, y}) // 加到队首
				}
			}
			pos[c] = nil // 避免重复使用传送门
		}

		// 下面代码和普通 BFS 是一样的
		for _, d := range dirs {
			x, y := p.x+d.x, p.y+d.y
			if 0 <= x && x < m && 0 <= y && y < n && matrix[x][y] != '#' && dis[p.x][p.y]+1 < dis[x][y] {
				dis[x][y] = dis[p.x][p.y] + 1
				q1 = append(q1, pair{x, y}) // 加到队尾
			}
		}
	}

	return -1
}
