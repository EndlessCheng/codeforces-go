package main

// github.com/EndlessCheng/codeforces-go
type pair struct{ x, y int }
var dir4 = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func wwork(n, m int, c [][]int) (ans int) {
	dis := make([][]int, n)
	filled := make([][]int, n) // 格子是否有人
	inQ := make([][]bool, n)
	for i := range dis {
		dis[i] = make([]int, m)
		filled[i] = make([]int, m)
		for j := range dis[i] {
			dis[i][j] = min(min(i, n-1-i), min(j, m-1-j))
			filled[i][j] = 1
		}
		inQ[i] = make([]bool, m)
	}

	// 倒序遍历
	for i := len(c) - 1; i >= 0; i-- {
		p := c[i]
		x, y := p[0]-1, p[1]-1
		ans += dis[x][y]
		filled[x][y] = 0
		q := []pair{{x, y}}
		for len(q) > 0 {
			p := q[0]
			q = q[1:]
			x, y := p.x, p.y
			inQ[x][y] = false
			for _, d := range dir4 {
				if xx, yy := x+d.x, y+d.y; 0 <= xx && xx < n && 0 <= yy && yy < m && dis[x][y]+filled[x][y] < dis[xx][yy] {
					dis[xx][yy] = dis[x][y] + filled[x][y]
					if !inQ[xx][yy] {
						inQ[xx][yy] = true
						q = append(q, pair{xx, yy})
					}
				}
			}
		}
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
