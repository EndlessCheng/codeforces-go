package main

import "math/bits"

// https://space.bilibili.com/206214
func minimumWeight(edges [][]int, queries [][]int) []int {
	n := len(edges) + 1
	type edge struct{ to, wt int }
	g := make([][]edge, n)
	for _, e := range edges {
		x, y, wt := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, wt})
		g[y] = append(g[y], edge{x, wt})
	}

	const mx = 17
	pa := make([][mx]int, n)
	dep := make([]int, n)
	dis := make([]int, n)
	var dfs func(int, int)
	dfs = func(x, p int) {
		pa[x][0] = p
		for _, e := range g[x] {
			y := e.to
			if y == p {
				continue
			}
			dep[y] = dep[x] + 1
			dis[y] = dis[x] + e.wt
			dfs(y, x)
		}
	}
	dfs(0, -1)

	for i := range mx - 1 {
		for x := range pa {
			p := pa[x][i]
			if p != -1 {
				pa[x][i+1] = pa[p][i]
			} else {
				pa[x][i+1] = -1
			}
		}
	}

	uptoDep := func(x, d int) int {
		for k := uint(dep[x] - d); k > 0; k &= k - 1 {
			x = pa[x][bits.TrailingZeros(k)]
		}
		return x
	}
	getLCA := func(x, y int) int {
		if dep[x] > dep[y] {
			x, y = y, x
		}
		y = uptoDep(y, dep[x])
		if y == x {
			return x
		}
		for i := mx - 1; i >= 0; i-- {
			if pv, pw := pa[x][i], pa[y][i]; pv != pw {
				x, y = pv, pw
			}
		}
		return pa[x][0]
	}
	getDis := func(x, y int) int { return dis[x] + dis[y] - dis[getLCA(x, y)]*2 }

	// 以上全是 LCA 模板

	ans := make([]int, len(queries))
	for i, q := range queries {
		a, b, c := q[0], q[1], q[2]
		ans[i] = (getDis(a, b) + getDis(b, c) + getDis(a, c)) / 2
	}
	return ans
}
