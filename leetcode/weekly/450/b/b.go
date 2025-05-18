package main

import (
	"cmp"
	"slices"
)

// https://space.bilibili.com/206214
// 完整的并查集模板，见我的数据结构题单
type unionFind struct {
	fa []int
	cc int // 连通块的个数
}

func newUnionFind(n int) unionFind {
	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	return unionFind{fa, n}
}

func (u unionFind) find(x int) int {
	if u.fa[x] != x {
		u.fa[x] = u.find(u.fa[x])
	}
	return u.fa[x]
}

func (u *unionFind) merge(from, to int) {
	x, y := u.find(from), u.find(to)
	if x == y {
		return
	}
	u.fa[x] = y
	u.cc-- // 合并后，连通块个数减一
}

func minSwaps(nums []int) int {
	n := len(nums)
	type tuple struct{ s, x, i int }
	a := make([]tuple, n)
	for i, num := range nums {
		s := 0
		for x := num; x > 0; x /= 10 {
			s += x % 10
		}
		a[i] = tuple{s, num, i}
	}

	slices.SortFunc(a, func(a, b tuple) int { return cmp.Or(a.s-b.s, a.x-b.x) })

	u := newUnionFind(n)
	for i, p := range a {
		u.merge(i, p.i)
	}
	return n - u.cc
}
