package main

// https://space.bilibili.com/206214
func magnificentSets(n int, edges [][]int) (ans int) {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0]-1, e[1]-1
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	time := make([]int, n) // 充当 vis 数组的作用（避免在 BFS 内部重复创建 vis 数组）
	clock := 0
	bfs := func(start int) (depth int) { // 返回从 start 出发的最大深度
		clock++
		time[start] = clock
		for q := []int{start}; len(q) > 0; depth++ {
			tmp := q
			q = nil
			for _, x := range tmp {
				for _, y := range g[x] {
					if time[y] != clock { // 没有在同一次 BFS 中访问过
						time[y] = clock
						q = append(q, y)
					}
				}
			}
		}
		return
	}

	colors := make([]int8, n)
	var nodes []int
	var isBipartite func(int, int8) bool
	isBipartite = func(x int, c int8) bool { // 二分图判定，原理见视频讲解
		nodes = append(nodes, x)
		colors[x] = c
		for _, y := range g[x] {
			if colors[y] == c || colors[y] == 0 && !isBipartite(y, -c) {
				return false
			}
		}
		return true
	}
	for i, c := range colors {
		if c == 0 {
			nodes = nil
			if !isBipartite(i, 1) { // 如果不是二分图（有奇环），则无法分组
				return -1
			}
			// 否则一定可以分组
			maxDepth := 0
			for _, x := range nodes { // 枚举连通块的每个点，作为起点 BFS，求最大深度
				maxDepth = max(maxDepth, bfs(x))
			}
			ans += maxDepth
		}
	}
	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
