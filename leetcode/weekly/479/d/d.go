package main

// https://space.bilibili.com/206214
func maxSubgraphScore1(n int, edges [][]int, good []int) []int {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	// subScore[x] 表示（以 0 为根时）包含 x 的子树 x 的最大得分（注意是子树不是子图）
	subScore := make([]int, n)
	// 计算 subScore[x]
	var dfs func(int, int)
	dfs = func(x, fa int) {
		subScore[x] = good[x]*2 - 1 // subScore[x] 一定包含 x
		for _, y := range g[x] {
			if y != fa {
				dfs(y, x)
				// 如果子树 y 的最大得分 > 0，选子树 y，否则不选
				subScore[x] += max(subScore[y], 0)
			}
		}
	}
	dfs(0, -1)

	ans := make([]int, n)
	ans[0] = subScore[0]
	// 对于 x 的儿子 y，计算包含 y 的子图最大得分
	var reroot func(int, int)
	reroot = func(x, fa int) {
		for _, y := range g[x] {
			if y != fa {
				// 从 ans[x] 中去掉子树 y。换根后，这部分内容变成 y 的一棵子树（记作 F）
				scoreF := ans[x] - max(subScore[y], 0)
				// 如果子树 F 的最大得分 > 0，选子树 F，否则不选
				ans[y] = subScore[y] + max(scoreF, 0)
				reroot(y, x)
			}
		}
	}
	reroot(0, -1)
	return ans
}

func maxSubgraphScore(n int, edges [][]int, ans []int) []int {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	var dfs func(int, int)
	dfs = func(x, fa int) {
		ans[x] = ans[x]*2 - 1
		for _, y := range g[x] {
			if y != fa {
				dfs(y, x)
				// 如果子树 y 的最大得分 > 0，选子树 y，否则不选
				ans[x] += max(ans[y], 0)
			}
		}
	}
	dfs(0, -1)

	// 对于 x 的儿子 y，计算包含 y 的子图最大得分
	var reroot func(int, int)
	reroot = func(x, fa int) {
		for _, y := range g[x] {
			if y != fa {
				// 从 ans[x] 中去掉子树 y。换根后，这部分内容变成 y 的一棵子树（记作 F）
				scoreF := ans[x] - max(ans[y], 0)
				// 如果子树 F 的最大得分 > 0，选子树 F，否则不选
				ans[y] += max(scoreF, 0)
				reroot(y, x)
			}
		}
	}
	reroot(0, -1)
	return ans
}
