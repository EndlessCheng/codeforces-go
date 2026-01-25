package main

import "slices"

// https://space.bilibili.com/206214
func specialNodes(n int, edges [][]int, x, y, z int) (ans int) {
	g := make([][]int, n)
	for _, e := range edges {
		v, w := e[0], e[1]
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	calcDis := func(start int) []int {
		dis := make([]int, n)
		var dfs func(int, int)
		dfs = func(v, fa int) {
			for _, w := range g[v] {
				if w != fa {
					dis[w] = dis[v] + 1
					dfs(w, v)
				}
			}
		}
		dfs(start, -1)
		return dis
	}

	dx := calcDis(x)
	dy := calcDis(y)
	dz := calcDis(z)

	for i := range n {
		a := []int{dx[i], dy[i], dz[i]}
		slices.Sort(a)
		if a[0]*a[0]+a[1]*a[1] == a[2]*a[2] {
			ans++
		}
	}
	return
}
