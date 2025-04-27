package main

// https://space.bilibili.com/206214
func baseUnitConversions(conversions [][]int) []int {
	const mod = 1_000_000_007
	n := len(conversions) + 1
	type edge struct{ to, weight int }
	g := make([][]edge, n)
	for _, e := range conversions {
		x, y, weight := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, weight})
	}

	ans := make([]int, n)
	var dfs func(int, int)
	dfs = func(x, mul int) {
		ans[x] = mul
		for _, e := range g[x] {
			dfs(e.to, mul*e.weight%mod)
		}
	}
	dfs(0, 1)
	return ans
}
