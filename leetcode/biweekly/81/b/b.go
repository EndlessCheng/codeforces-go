package main

// https://space.bilibili.com/206214/dynamic
func countPairs(n int, edges [][]int) (ans int64) {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	vis := make([]bool, n)
	tot, size := 0, 0
	var dfs func(int)
	dfs = func(x int) {
		vis[x] = true
		size++
		for _, y := range g[x] {
			if !vis[y] {
				dfs(y)
			}
		}
	}
	for i, b := range vis {
		if !b {
			size = 0
			dfs(i)
			ans += int64(size) * int64(tot)
			tot += size
		}
	}
	return
}
