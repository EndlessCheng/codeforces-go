package main

// https://space.bilibili.com/206214
func assignEdgeWeights(edges [][]int) int {
	n := len(edges) + 1
	g := make([][]int, n+1)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	var dfs func(int, int) int
	dfs = func(x, fa int) (d int) {
		for _, y := range g[x] {
			if y != fa {
				d = max(d, dfs(y, x))
			}
		}
		return d + 1
	}

	k := dfs(1, 0) - 1 // dfs 返回的是节点个数，减一得到边的个数
	return pow(2, k-1)
}

func pow(x, n int) int {
	const mod = 1_000_000_007
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}
