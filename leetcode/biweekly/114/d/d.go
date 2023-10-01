package main

// https://space.bilibili.com/206214
func maxKDivisibleComponents(n int, edges [][]int, values []int, k int) (ans int) {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	var dfs func(int, int) int
	dfs = func(x, fa int) int {
		s := values[x]
		for _, y := range g[x] {
			if y != fa {
				s += dfs(y, x)
			}
		}
		if s%k == 0 {
			ans++
		}
		return s
	}
	dfs(0, -1)
	return
}
