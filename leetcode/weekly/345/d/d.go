package main

// https://space.bilibili.com/206214
func countCompleteComponents(n int, edges [][]int) (ans int) {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	vis := make([]bool, n)
	var v, e int
	var dfs func(int)
	dfs = func(x int) {
		vis[x] = true
		v++
		e += len(g[x])
		for _, y := range g[x] {
			if !vis[y] {
				dfs(y)
			}
		}
	}
	for i, b := range vis {
		if !b {
			v, e = 0, 0
			dfs(i)
			if e == v*(v-1) {
				ans++
			}
		}
	}
	return
}
