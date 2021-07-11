package main

// github.com/EndlessCheng/codeforces-go
type pair struct{ x, y int }

var dir4 = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func nearestExit(g [][]byte, entrance []int) int {
	n, m := len(g), len(g[0])
	s := pair{entrance[0], entrance[1]}
	g[s.x][s.y] = 0
	q := []pair{s}
	for ans := 1; len(q) > 0; ans++ {
		tmp := q
		q = nil
		for _, p := range tmp {
			for _, d := range dir4 {
				if x, y := p.x+d.x, p.y+d.y; 0 <= x && x < n && 0 <= y && y < m && g[x][y] == '.' {
					if x == 0 || y == 0 || x == n-1 || y == m-1 {
						return ans
					}
					g[x][y] = 0
					q = append(q, pair{x, y})
				}
			}
		}
	}
	return -1
}
