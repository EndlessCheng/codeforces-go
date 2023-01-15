package main

// https://space.bilibili.com/206214
func maxOutput(n int, edges [][]int, price []int) int64 {
	ans := 0
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x) // 建树
	}
	// 返回带叶子的最大路径和，不带叶子的最大路径和
	var dfs func(int, int) (int, int)
	dfs = func(x, fa int) (int, int) {
		p := price[x]
		maxS1, maxS2 := p, 0
		for _, y := range g[x] {
			if y != fa {
				s1, s2 := dfs(y, x)
				// 前面最大带叶子的路径 + 当前不带叶子的路径
				// 前面最大不带叶子的路径 + 当前带叶子的路径
				ans = max(ans, max(maxS1+s2, maxS2+s1))
				maxS1 = max(maxS1, s1+p)
				maxS2 = max(maxS2, s2+p) // 这里加上 p 是因为 x 必然不是叶子
			}
		}
		return maxS1, maxS2
	}
	dfs(0, -1)
	return int64(ans)
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
