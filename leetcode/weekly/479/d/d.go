package main

// https://space.bilibili.com/206214
func maxSubgraphScore(n int, edges [][]int, good []int) []int {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	// subScore[x] 表示（以 0 为根时）子树 x 的最大得分（一定包含节点 x）
	subScore := make([]int, n)
	// 计算并返回 subScore[x]
	var dfs func(int, int) int
	dfs = func(x, fa int) int {
		for _, y := range g[x] {
			if y != fa {
				// 如果子树 y 的得分是负数，不选子树 y，否则选子树 y
				subScore[x] += max(dfs(y, x), 0)
			}
		}
		subScore[x] += good[x]*2 - 1 // subScore[x] 一定包含 x
		return subScore[x]
	}
	dfs(0, -1)

	ans := make([]int, n)
	// 计算子图 x 的最大得分 scoreX，其中 faScore 表示来自父节点 fa 的最大得分（一定包含节点 fa）
	var reroot func(int, int, int)
	reroot = func(x, fa, faScore int) {
		scoreX := subScore[x] + max(faScore, 0)
		ans[x] = scoreX
		for _, y := range g[x] {
			if y != fa {
				// scoreX-max(subScore[y],0) 是不含子树 y 的最大得分
				reroot(y, x, scoreX-max(subScore[y], 0))
			}
		}
	}
	reroot(0, -1, 0)
	return ans
}
