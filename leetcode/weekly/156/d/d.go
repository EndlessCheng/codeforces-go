package main

// https://space.bilibili.com/206214
type tuple struct{ x, y, s int }
var dirs = []tuple{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}

func minimumMoves(g [][]int) int {
	n := len(g)
	vis := make([][][2]bool, n)
	for i := range vis {
		vis[i] = make([][2]bool, n)
	}
	vis[0][0][0] = true
	q := []tuple{{}}
	for step := 1; len(q) > 0; step++ {
		tmp := q
		q = nil
		for _, t := range tmp {
			for _, d := range dirs {
				x, y, s := t.x+d.x, t.y+d.y, t.s^d.s
				x2, y2 := x+s, y+(s^1) // 蛇头
				if x2 < n && y2 < n && !vis[x][y][s] &&
					g[x][y] == 0 && g[x2][y2] == 0 && (d.s == 0 || g[x+1][y+1] == 0) {
					if x == n-1 && y == n-2 { // 此时蛇头一定在 (n-1,n-1)
						return step
					}
					vis[x][y][s] = true
					q = append(q, tuple{x, y, s})
				}
			}
		}
	}
	return -1
}
