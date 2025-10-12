package main

// https://space.bilibili.com/206214

// 785. 判断二分图
func isBipartite(graph [][]int) bool {
	// colors[i] = 0  表示未访问节点 i
	// colors[i] = 1  表示节点 i 为红色
	// colors[i] = -1 表示节点 i 为蓝色
	colors := make([]int8, len(graph))

	var dfs func(int, int8) bool
	dfs = func(x int, c int8) bool {
		colors[x] = c // 节点 x 染成颜色 c
		for _, y := range graph[x] {
			// 邻居 y 的颜色与 x 的相同，说明不是二分图，返回 false
			// 或者继续递归，发现不是二分图，返回 false
			if colors[y] == c ||
				colors[y] == 0 && !dfs(y, -c) { // 取相反数，实现交替染色
				return false
			}
		}
		return true
	}

	// 可能有多个连通块
	for i, c := range colors {
		if c == 0 && !dfs(i, 1) {
			// 从节点 i 开始递归，发现 i 所在连通块不是二分图
			return false
		}
	}
	return true
}

func possibleBipartition(n int, dislikes [][]int) bool {
	g := make([][]int, n)
	for _, e := range dislikes {
		x, y := e[0]-1, e[1]-1 // 节点编号改成从 0 开始
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	return isBipartite(g)
}
