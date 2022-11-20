package main

// https://space.bilibili.com/206214
func minimumFuelCost(roads [][]int, seats int) int64 {
	ans := 0
	g := make([][]int, len(roads)+1)
	for _, e := range roads {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	var dfs func(int, int) int
	dfs = func(x, fa int) int {
		size := 1
		for _, y := range g[x] {
			if y != fa {
				size += dfs(y, x)
			}
		}
		if x > 0 {
			ans += (size + seats - 1) / seats
		}
		return size
	}
	dfs(0, -1)
	return int64(ans)
}
