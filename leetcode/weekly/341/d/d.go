package main

// https://space.bilibili.com/206214
func minimumTotalPrice(n int, edges [][]int, price []int, trips [][]int) int {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x) // 建树
	}

	qs := make([][]int, n)
	for _, t := range trips {
		x, y := t[0], t[1]
		qs[x] = append(qs[x], y) // 路径端点分组
		if x != y {
			qs[y] = append(qs[y], x)
		}
	}

	// 并查集模板
	pa := make([]int, n)
	for i := range pa {
		pa[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if pa[x] != x {
			pa[x] = find(pa[x])
		}
		return pa[x]
	}

	diff := make([]int, n)
	father := make([]int, n)
	color := make([]int8, n)
	var tarjan func(int, int)
	tarjan = func(x, fa int) {
		father[x] = fa
		color[x] = 1 // 递归中
		for _, y := range g[x] {
			if color[y] == 0 { // 未递归
				tarjan(y, x)
				pa[y] = x // 相当于把 y 的子树节点全部 merge 到 x
			}
		}
		for _, y := range qs[x] {
			// color[y] == 2 意味着 y 所在子树已经遍历完
			// 也就意味着 y 已经 merge 到它和 x 的 lca 上了
			if y == x || color[y] == 2 { // 从 y 向上到达 lca 然后拐弯向下到达 x
				diff[x]++
				diff[y]++
				lca := find(y)
				diff[lca]--
				if f := father[lca]; f >= 0 {
					diff[f]--
				}
			}
		}
		color[x] = 2 // 递归结束
	}
	tarjan(0, -1)

	var dfs func(int, int) (int, int, int)
	dfs = func(x, fa int) (notHalve, halve, cnt int) {
		cnt = diff[x]
		for _, y := range g[x] {
			if y != fa {
				nh, h, c := dfs(y, x)  // 计算 y 不变/减半的最小价值总和
				notHalve += min(nh, h) // x 不变，那么 y 可以不变，可以减半，取这两种情况的最小值
				halve += nh            // x 减半，那么 y 只能不变
				cnt += c               // 自底向上累加差分值
			}
		}
		notHalve += price[x] * cnt  // x 不变
		halve += price[x] * cnt / 2 // x 减半
		return
	}
	nh, h, _ := dfs(0, -1)
	return min(nh, h)
}
