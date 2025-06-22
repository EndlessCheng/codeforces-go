package main

// https://space.bilibili.com/206214
func minIncrease(n int, edges [][]int, cost []int) (ans int) {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	g[0] = append(g[0], -1) // 避免误把根节点当作叶子

	var dfs func(int, int, int) int
	dfs = func(x, fa, pathSum int) (maxS int) {
		pathSum += cost[x]
		if len(g[x]) == 1 {
			return pathSum
		}

		cnt := 0 // 在根到叶子的 pathSum 中，有 cnt 个 pathSum 等于 maxS
		for _, y := range g[x] {
			if y == fa {
				continue
			}
			mx := dfs(y, x, pathSum)
			if mx > maxS {
				maxS = mx
				cnt = 1
			} else if mx == maxS {
				cnt++
			}
		}
		// 其余小于 maxS 的 pathSum，可以通过增大 cost[y] 的值，改成 maxS
		ans += len(g[x]) - 1 - cnt
		return maxS
	}
	dfs(0, -1, 0)
	return
}
