package main

// github.com/EndlessCheng/codeforces-go
func friendRequests(n int, restrictions [][]int, requests [][]int) []bool {
	fa := make([]int, n) // 初始化并查集
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

	cant := make([]map[int]bool, n)
	for i := range cant {
		cant[i] = map[int]bool{}
	}
	for _, r := range restrictions { // 初始化无法成为朋友的用户
		v, w := r[0], r[1]
		cant[v][w] = true
		cant[w][v] = true // 双向
	}

	ans := make([]bool, len(requests))
	for i, r := range requests {
		v, w := find(r[0]), find(r[1])
		if v == w { // 已经是直接朋友
			ans[i] = true
			continue
		}
		if cant[v][w] { // 无法成为朋友
			continue
		}
		ans[i] = true
		if len(cant[v]) > len(cant[w]) { // 常数优化：总是从小的集合合并到大的集合上
			v, w = w, v
		}
		for x := range cant[v] { // 将 cant[v] 合并到 cant[w] 上
			x = find(x)
			cant[w][x] = true
			cant[x][w] = true // 双向
		}
		fa[v] = w
	}
	return ans
}
