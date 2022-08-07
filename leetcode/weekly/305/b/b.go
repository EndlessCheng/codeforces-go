package main

// https://space.bilibili.com/206214
func reachableNodes(n int, edges [][]int, restricted []int) (ans int) {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	r := map[int]bool{}
	for _, x := range restricted {
		r[x] = true
	}
	var dfs func(int, int)
	dfs = func(x, fa int) {
		ans++
		for _, y := range g[x] {
			if y != fa && !r[y] {
				dfs(y, x)
			}
		}
	}
	dfs(0, -1)
	return
}
