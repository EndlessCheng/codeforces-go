package main

import "slices"

// https://space.bilibili.com/206214

// 模板来源 https://leetcode.cn/circle/discuss/mOr1u6/
type unionFind struct {
	fa  []int // 代表元
	odd []int // 集合中的奇数个数
}

func newUnionFind(n int) unionFind {
	fa := make([]int, n)
	odd := make([]int, n)
	// 一开始有 n 个集合 {0}, {1}, ..., {n-1}
	// 集合 i 的代表元是自己
	for i := range fa {
		fa[i] = i
		odd[i] = i % 2
	}
	return unionFind{fa, odd}
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
	u.fa[x] = y          // 合并集合
	u.odd[y] += u.odd[x] // 更新集合中的奇数个数
}

func maxAlternatingSum(nums []int, swaps [][]int) (ans int64) {
	n := len(nums)
	uf := newUnionFind(n)
	for _, p := range swaps {
		uf.merge(p[0], p[1])
	}

	g := make([][]int, n)
	for i, x := range nums {
		f := uf.find(i)
		g[f] = append(g[f], x) // 相同集合的元素分到同一组
	}

	for i, a := range g {
		if a == nil {
			continue
		}
		slices.Sort(a)
		odd := uf.odd[i]
		// 小的取负号，大的取正号
		for j, x := range a {
			if j < odd {
				ans -= int64(x)
			} else {
				ans += int64(x)
			}
		}
	}
	return
}
