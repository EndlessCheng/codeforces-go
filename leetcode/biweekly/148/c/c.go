package main

// https://space.bilibili.com/206214
func longestSpecialPath(edges [][]int, nums []int) []int {
	type edge struct{ to, weight int }
	g := make([][]edge, len(nums))
	for _, e := range edges {
		x, y, w := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, w})
		g[y] = append(g[y], edge{x, w})
	}

	maxLen := -1
	minNodes := 0
	dis := []int{0}
	// 颜色 -> 该颜色最近一次出现的深度 +1，注意这里已经 +1 了
	lastDepth := map[int]int{}

	var dfs func(int, int, int)
	dfs = func(x, fa, topDepth int) {
		color := nums[x]
		oldDepth := lastDepth[color]
		topDepth = max(topDepth, oldDepth)

		length := dis[len(dis)-1] - dis[topDepth]
		nodes := len(dis) - topDepth
		if length > maxLen || length == maxLen && nodes < minNodes {
			maxLen = length
			minNodes = nodes
		}

		lastDepth[color] = len(dis)
		for _, e := range g[x] {
			y := e.to
			if y != fa { // 避免访问父节点
				dis = append(dis, dis[len(dis)-1]+e.weight)
				dfs(y, x, topDepth)
				dis = dis[:len(dis)-1] // 恢复现场
			}
		}
		lastDepth[color] = oldDepth // 恢复现场
	}

	dfs(0, -1, 0)
	return []int{maxLen, minNodes}
}
