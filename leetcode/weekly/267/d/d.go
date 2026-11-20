package main

// github.com/EndlessCheng/codeforces-go
func friendRequests1(n int, restrictions [][]int, requests [][]int) []bool {
	// 初始化并查集
	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	ans := make([]bool, len(requests))
next:
	for i, req := range requests {
		u, v := find(req[0]), find(req[1])
		// 暴力检查
		for _, r := range restrictions {
			x, y := find(r[0]), find(r[1])
			if x == u && y == v || x == v && y == u {
				continue next // 不能连边
			}
		}
		// 可以连边
		fa[u] = v // 并查集合并
		ans[i] = true
	}
	return ans
}

func friendRequests(n int, restrictions [][]int, requests [][]int) []bool {
	// 初始化并查集
	fa := make([]int, n)
	size := make([]int, n)
	for i := range fa {
		fa[i] = i
		size[i] = 1
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	// ban[i][x] 表示代表元为 i 的连通块不能与代表元为 x 的连通块合并
	ban := make([]map[int]bool, n)
	for i := range ban {
		ban[i] = map[int]bool{}
	}
	// 初始化无法合并的连通块
	for _, r := range restrictions {
		x, y := r[0], r[1]
		ban[x][y] = true
		ban[y][x] = true // 双向
	}

	ans := make([]bool, len(requests))
	for i, r := range requests {
		// 把代表元为 u 的连通块与代表元为 v 的连通块合并
		u, v := find(r[0]), find(r[1])
		if u == v { // 已经合并
			ans[i] = true
			continue
		}
		if ban[u][v] { // 无法合并
			continue
		}
		// 启发式合并：总是把小的连通块合并到大的连通块
		if size[u] > size[v] {
			u, v = v, u
		}
		// 把 u 并入 v，那么无法与 u 合并的连通块，也无法与 v 合并
		for x := range ban[u] {
			ban[v][x] = true
			delete(ban[x], u) // 删掉旧的
			ban[x][v] = true // 添加新的
		}
		ban[u] = nil // 保证空间复杂度
		fa[u] = v // 并查集合并
		size[v] += size[u]
		ans[i] = true
	}
	return ans
}
