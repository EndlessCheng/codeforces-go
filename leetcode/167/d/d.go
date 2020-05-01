package main

func shortestPath(g [][]int, kk int) int {
	n, m := int8(len(g)), int8(len(g[0]))
	if kk > int(n+m-3) {
		kk = int(n + m - 3)
	}
	if kk < 0 {
		kk = 0
	}
	K := int8(kk)
	vis := [40][40][78]bool{}
	vis[n-1][m-1][K] = true
	type pair struct {
		x, y, k int8
		dep     int16
	}
	dir4 := [...][2]int8{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右
	for q := []pair{{n - 1, m - 1, K, 0}}; len(q) > 0; {
		p := q[0]
		q = q[1:]
		if p.x == 0 && p.y == 0 {
			return int(p.dep)
		}
		for _, d := range dir4 {
			if xx, yy := p.x+d[0], p.y+d[1]; xx >= 0 && xx < n && yy >= 0 && yy < m {
				if g[xx][yy] == 0 {
					if !vis[xx][yy][p.k] {
						vis[xx][yy][p.k] = true
						q = append(q, pair{xx, yy, p.k, p.dep + 1})
					}
				} else {
					if p.k > 0 && !vis[xx][yy][p.k-1] {
						vis[xx][yy][p.k-1] = true
						q = append(q, pair{xx, yy, p.k - 1, p.dep + 1})
					}
				}
			}
		}
	}
	return -1
}
