package main

// https://space.bilibili.com/206214
func minCost(n int, colors string, edges [][]int) int64 {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	base, mx := 0, 0

	var dfs func(int, int) (int, int)
	dfs = func(x, fa int) (int, int) {
		size := 1
		f := 0
		// 遍历 x 的子树
		for _, y := range g[x] {
			if y != fa {
				sz, fy := dfs(y, x)
				size += sz
				// 如果 fy < 0，那么子树 y 都变成蓝色是最优的，这已经算入 base 中
				f += max(fy, 0)
			}
		}

		// 考虑节点 x
		if colors[x] == 'R' {
			base += size // 把 x 变成蓝色的代价
			f += size // 把 x 恢复成红色可以让 base 减少 size
		} else {
			f -= size // 把 x 变成红色会让 base 增大 size
		}
		mx = max(mx, f) // 最多可以让 base 减少多少

		return size, f
	}

	dfs(0, -1)
	return int64(base - mx)
}
