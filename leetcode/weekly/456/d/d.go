package main

import (
	"math"
	"slices"
	"sort"
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

	left := uf.cc - 1
	if left == 0 { // 只需选必选边
		return minS1
	}

	ans := minS1
	// Kruskal 算法求最大生成树
	slices.SortFunc(edges, func(a, b []int) int { return b[2] - a[2] })
	for _, e := range edges {
		x, y, s, must := e[0], e[1], e[2], e[3]
		if must == 0 && uf.merge(x, y) {
			if left > k {
				ans = min(ans, s)
			} else {
				ans = min(ans, s*2)
			}
			left--
			if left == 0 { // 已经得到生成树了
				break
			}
		}
	}
	return ans
}

func maxStability1(n int, edges [][]int, k int) int {
	mustUf := newUnionFind(n) // 必选边并查集
	allUf := newUnionFind(n)  // 全图并查集
	minS, maxS := math.MaxInt, 0
	for _, e := range edges {
		x, y, s, must := e[0], e[1], e[2], e[3]
		if must > 0 && !mustUf.merge(x, y) { // 必选边成环
			return -1
		}
		allUf.merge(x, y)
		minS = min(minS, s)
		maxS = max(maxS, s)
	}

	if allUf.cc > 1 { // 图不连通
		return -1
	}

	left, right := minS, maxS*2
	ans := left + sort.Search(right-left, func(low int) bool {
		low += left
		low++ // 二分最小的不满足要求的 low+1，那么答案就是最大的满足要求的 low
		u := newUnionFind(n)
		for _, e := range edges {
			x, y, s, must := e[0], e[1], e[2], e[3]
			if must > 0 && s < low { // 必选边的边权太小
				return true
			}
			if must > 0 || s >= low {
				u.merge(x, y)
			}
		}

		leftK := k
		for _, e := range edges {
			if leftK == 0 || u.cc == 1 {
				break
			}
			x, y, s, must := e[0], e[1], e[2], e[3]
			if must == 0 && s < low && s*2 >= low && u.merge(x, y) {
				leftK--
			}
		}
		return u.cc > 1
	})
	return ans
}
