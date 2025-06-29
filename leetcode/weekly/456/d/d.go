package main

import (
	"math"
	"slices"
)

// https://space.bilibili.com/206214
type unionFind struct {
	fa []int // 代表元
	cc int   // 连通块个数
}

func newUnionFind(n int) unionFind {
	fa := make([]int, n)
	// 一开始有 n 个集合 {0}, {1}, ..., {n-1}
	// 集合 i 的代表元是自己
	for i := range fa {
		fa[i] = i
	}
	return unionFind{fa, n}
}

// 返回 x 所在集合的代表元
// 同时做路径压缩，也就是把 x 所在集合中的所有元素的 fa 都改成代表元
func (u unionFind) find(x int) int {
	// 如果 fa[x] == x，则表示 x 是代表元
	if u.fa[x] != x {
		u.fa[x] = u.find(u.fa[x]) // fa 改成代表元
	}
	return u.fa[x]
}

// 把 from 所在集合合并到 to 所在集合中
// 返回是否合并成功
func (u *unionFind) merge(from, to int) bool {
	x, y := u.find(from), u.find(to)
	if x == y { // from 和 to 在同一个集合，不做合并
		return false
	}
	u.fa[x] = y // 合并集合。修改后就可以认为 from 和 to 在同一个集合了
	u.cc--      // 成功合并，连通块个数减一
	return true
}

func maxStability(n int, edges [][]int, k int) int {
	if len(edges) < n-1 { // 图不连通
		return -1
	}

	uf := newUnionFind(n)
	allUf := newUnionFind(n)
	minS1 := math.MaxInt
	for _, e := range edges {
		x, y, s, must := e[0], e[1], e[2], e[3]
		if must > 0 {
			if !uf.merge(x, y) { // 必选边成环
				return -1
			}
			minS1 = min(minS1, s)
		}
		allUf.merge(x, y)
	}

	if allUf.cc > 1 { // 图不连通
		return -1
	}

	// Kruskal 算法求最大生成树
	slices.SortFunc(edges, func(a, b []int) int { return b[2] - a[2] })
	a := []int{}
	for _, e := range edges {
		x, y, s, must := e[0], e[1], e[2], e[3]
		if must == 0 && uf.merge(x, y) {
			a = append(a, s)
		}
	}

	// 如下三者的最小值：
	// 1. must = 1 中的最小值
	// 2. a 中的最小边权 * 2
	// 3. a 中的第 k+1 小边权
	m := len(a)
	if m == 0 {
		return minS1
	}
	ans := min(minS1, a[m-1]*2)
	if k < m {
		ans = min(ans, a[m-1-k])
	}
	return ans
}
