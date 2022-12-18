package main

// https://space.bilibili.com/206214
func isPossible(n int, edges [][]int) bool {
	g := map[int]map[int]bool{}
	for _, e := range edges {
		x, y := e[0], e[1]
		if g[x] == nil {
			g[x] = map[int]bool{}
		}
		g[x][y] = true
		if g[y] == nil {
			g[y] = map[int]bool{}
		}
		g[y][x] = true
	}
	odd := []int{}
	for i, nb := range g {
		if len(nb)%2 > 0 {
			odd = append(odd, i)
		}
	}
	m := len(odd)
	if m == 0 {
		return true
	}
	if m == 2 {
		x, y := odd[0], odd[1]
		if !g[x][y] {
			return true
		}
		for i := 1; i <= n; i++ {
			if i != x && i != y && !g[i][x] && !g[i][y] {
				return true
			}
		}
		return false
	}
	if m == 4 {
		a, b, c, d := odd[0], odd[1], odd[2], odd[3]
		return !g[a][b] && !g[c][d] || !g[a][c] && !g[b][d] || !g[a][d] && !g[b][c]
	}
	return false
}
