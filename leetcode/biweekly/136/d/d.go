package main

// https://space.bilibili.com/206214
func timeTaken(edges [][]int) []int {
	n := len(edges) + 1
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	// subRes[x] 保存子树 x 的最大深度 maxD，次大深度 maxD2，以及最大深度要往儿子 y 走
	subRes := make([]struct{ maxD, maxD2, y int }, n)
	// 计算 subRes[x]
	var dfs func(int, int)
	dfs = func(x, fa int) {
		res := &subRes[x]
		for _, y := range g[x] {
			if y == fa {
				continue
			}
			dfs(y, x)
			w := 2 - y%2 // 从 x 到 y 的边权
			maxD := subRes[y].maxD + w // 从 x 出发，往 y 方向的最大深度
			if maxD > res.maxD {
				res.maxD2 = res.maxD
				res.maxD = maxD
				res.y = y
			} else if maxD > res.maxD2 {
				res.maxD2 = maxD
			}
		}
	}
	dfs(0, -1)

	// ans[x] 表示当 x 是树根时，整棵树的最大深度
	ans := make([]int, n)
	// 计算 ans[x]
	var reroot func(int, int, int)
	reroot = func(x, fa, fromUp int) {
		p := subRes[x]
		ans[x] = max(subRes[x].maxD, fromUp)
		for _, y := range g[x] {
			if y == fa {
				continue
			}
			// 站在 x 的角度，不往 y 走，能走多远？
			// 要么往上走（fromUp），要么往除了 y 的其余子树走（mx），二者取最大值
			mx := p.maxD
			if y == p.y { // 对于 y 来说，上面要选次大的
				mx = p.maxD2
			}
			w := 2 - x%2 // 从 y 到 x 的边权
			reroot(y, x, max(fromUp, mx)+w) // 对于 y 来说，加上从 y 到 x 的边权
		}
	}
	reroot(0, -1, 0)
	return ans
}
