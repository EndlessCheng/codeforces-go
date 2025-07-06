package main

import "slices"

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
func (u *unionFind) merge(from, to int) {
	x, y := u.find(from), u.find(to)
	if x == y { // from 和 to 在同一个集合，不做合并
		return
	}
	u.fa[x] = y // 合并集合。修改后就可以认为 from 和 to 在同一个集合了
	u.cc--      // 成功合并，连通块个数减一
}

func minTime(n int, edges [][]int, k int) int {
	slices.SortFunc(edges, func(a, b []int) int { return b[2] - a[2] })

	u := newUnionFind(n)
	for _, e := range edges {
		u.merge(e[0], e[1])
		if u.cc < k { // 这条边不能留，即移除所有 time <= e[2] 的边
			return e[2]
		}
	}
	return 0 // 无需移除任何边
}
