package main

import "math/bits"

// https://space.bilibili.com/206214
func findMedian(n int, edges [][]int, queries [][]int) []int {
	type edge struct{ to, wt int }
	g := make([][]edge, n)
	for _, e := range edges {
		x, y, wt := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, wt})
		g[y] = append(g[y], edge{x, wt})
	}

	const mx = 17 // 也可以写 bits.Len(uint(n))
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

	// 返回 x 和 y 的最近公共祖先（节点编号从 0 开始）
	getLCA := func(x, y int) int {
		if dep[x] > dep[y] {
			x, y = y, x
		}
		y = uptoDep(y, dep[x]) // 使 y 和 x 在同一深度
		if y == x {
			return x
		}
		for i := mx - 1; i >= 0; i-- {
			px, py := pa[x][i], pa[y][i]
			if px != py {
				x, y = px, py // 同时往上跳 2^i 步
			}
		}
		return pa[x][0]
	}

	// 从 x 往上跳【至多】d 距离，返回最远能到达的节点 
	uptoDis := func(x, d int) int {
		dx := dis[x]
		for i := mx - 1; i >= 0; i-- {
			p := pa[x][i]
			if p != -1 && dx-dis[p] <= d { // 可以跳至多 d
				x = p
			}
		}
		return x
	}

	// 以上是 LCA 模板

	ans := make([]int, len(queries))
	for i, q := range queries {
		x, y := q[0], q[1]
		if x == y {
			ans[i] = x
			continue
		}
		lca := getLCA(x, y)
		disXY := dis[x] + dis[y] - dis[lca]*2
		half := (disXY + 1) / 2
		if dis[x]-dis[lca] >= half { // 答案在 x-lca 路径中
			// 先往上跳至多 half-1，然后再跳一步，就是至少 half
			to := uptoDis(x, half-1)
			ans[i] = pa[to][0] // 再跳一步
		} else { // 答案在 y-lca 路径中
			// 从 y 出发至多 disXY-half，就是从 x 出发至少 half
			ans[i] = uptoDis(y, disXY-half) 
		}
	}
	return ans
}
