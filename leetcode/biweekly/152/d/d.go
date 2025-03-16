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
	// 颜色 -> 该颜色最近一次出现的深度 +1，注意这里已经 +1 了，下面不需要再 +1
	lastDepth := map[int]int{}

	var dfs func(int, int, int, int)
	dfs = func(x, fa, topDepth, last1 int) {
		color := nums[x]
		last2 := lastDepth[color]
		topDepth = max(topDepth, min(last1, last2)) // 维护窗口左端点

		length := dis[len(dis)-1] - dis[topDepth]
		nodes := len(dis) - topDepth
		if length > maxLen || length == maxLen && nodes < minNodes {
			maxLen = length
			minNodes = nodes
		}

		lastDepth[color] = len(dis)
		for _, e := range g[x] {
			y := e.to
			if y != fa {
				dis = append(dis, dis[len(dis)-1]+e.weight)
				dfs(y, x, topDepth, max(last1, last2)) // 维护 last1
				dis = dis[:len(dis)-1] // 恢复现场
			}
		}
		lastDepth[color] = last2 // 恢复现场
	}

	dfs(0, -1, 0, 0)
	return []int{maxLen, minNodes}
}
