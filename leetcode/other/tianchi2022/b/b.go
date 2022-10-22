package main

// https://space.bilibili.com/206214
var dir4 = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func getLength(g []string) (ans int) {
	n, m := len(g), len(g[0])
	x, y, di := 0, 0, 1
	for 0 <= x && x < n && 0 <= y && y < m {
		ans++
		c := g[x][y]
		if c == 'L' {
			di ^= 2
		} else if c == 'R' {
			di ^= 3
		}
		x += dir4[di].x
		y += dir4[di].y
	}
	return
}
