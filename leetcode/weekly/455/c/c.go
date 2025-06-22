package main

// https://space.bilibili.com/206214
func minIncrease(n int, edges [][]int, cost []int) (ans int) {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	g[0] = append(g[0], -1)

	var dfs func(int, int) int
	dfs = func(x, fa int) (maxS int) {
		cnt := 0
		for _, y := range g[x] {
			if y == fa {
				continue
			}
			mx := dfs(y, x)
			if mx > maxS {
				maxS = mx
				cnt = 1
			} else if mx == maxS {
				cnt++
			}
		}
		ans += len(g[x]) - 1 - cnt
		return maxS + cost[x]
	}
	dfs(0, -1)
	return
}
