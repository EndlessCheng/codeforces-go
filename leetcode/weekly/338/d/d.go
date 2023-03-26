package main

// https://space.bilibili.com/206214
func collectTheCoins(coins []int, edges [][]int) int {
	n := len(coins)
	g := make([][]int, n)
	deg := make([]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x) // 建图
		deg[x]++
		deg[y]++
	}

	ans := n - 1

	// 用拓扑排序「剪枝」：去掉没有金币的子树
	q := make([]int, 0, n)
	for i, d := range deg {
		if d == 1 && coins[i] == 0 { // 无金币叶子
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		ans--
		deg[x] = 0
		for _, y := range g[x] {
			deg[y]--
			if deg[y] == 1 && coins[y] == 0 {
				q = append(q, y)
			}
		}
	}

	// 再次拓扑排序
	for i, d := range deg {
		if d == 1 { // 有金币叶子
			q = append(q, i)
		}
	}
	ans -= len(q)
	for _, x := range q {
		for _, y := range g[x] {
			deg[y]--
			if deg[y] == 1 {
				ans--
			}
		}
	}
	return max(ans * 2, 0)
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
