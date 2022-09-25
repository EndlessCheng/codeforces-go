package main

import "sort"

// https://space.bilibili.com/206214
func numberOfGoodPaths(vals []int, edges [][]int) int {
	n := len(vals)
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x) // 建图
	}

	// 并查集模板
	fa := make([]int, n)
	size := make([]int, n) // size[x] 表示节点值等于 vals[x] 的节点个数
	id := make([]int, n)   // 后面排序用
	for i := range fa {
		fa[i] = i
		size[i] = 1
		id[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	ans := n
	sort.Slice(id, func(i, j int) bool { return vals[id[i]] < vals[id[j]] })
	for _, x := range id {
		vx := vals[x]
		fx := find(x)
		for _, y := range g[x] {
			y = find(y)
			if y == fx || vals[y] > vx { // 只考虑最大节点值比 vx 小的连通块
				continue
			}
			if vals[y] == vx { // 可以构成好路径
				ans += size[fx] * size[y] // 乘法原理
				size[fx] += size[y]       // 统计连通块内节点值等于 vx 的节点个数
			}
			fa[y] = fx // 关键：把小的节点值合并到大的节点值上
		}
	}
	return ans
}
