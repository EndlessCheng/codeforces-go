package main

// https://space.bilibili.com/206214
type pair struct{ x, y, d int }
var dirs = []pair{{0, 1, 0}, {1, 0, 0}, {0, 0, 1}}

func minimumMoves(g [][]int) int {
	n := len(g)
	dis := make([][][2]int, n)
	for i := range dis {
		dis[i] = make([][2]int, n)
		for j := range dis[i] {
			dis[i][j] = [2]int{-1, -1}
		}
	}
	dis[0][0][0] = 0
	q := []pair{{}}
	for ans := 1; len(q) > 0; ans++ {
		tmp := q
		q = nil
		for _, p := range tmp {
			for _, dir := range dirs {
				x, y, d := p.x+dir.x, p.y+dir.y, p.d^dir.d
				x2, y2 := d+x, d^1+y
				if x2 < n && y2 < n && dis[x][y][d] < 0 && g[x][y] == 0 && g[x2][y2] == 0 && (dir.d == 0 || g[x+1][y+1] == 0) {
					if x == n-1 && y == n-2 {
						return ans
					}
					dis[x][y][d] = ans
					q = append(q, pair{x, y, d})
				}
			}
		}
	}
	return -1
}
