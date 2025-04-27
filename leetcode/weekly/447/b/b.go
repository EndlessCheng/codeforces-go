package main

import "sort"

// https://space.bilibili.com/206214
func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []bool {
	// 每个节点所在连通块的编号
	id := make([]int, n)
	for i := 1; i < n; i++ {
		id[i] = id[i-1]
		if nums[i]-nums[i-1] > maxDiff {
			id[i]++ // 创建一个新的连通块
		}
	}

	ans := make([]bool, len(queries))
	for i, q := range queries {
		ans[i] = id[q[0]] == id[q[1]]
	}
	return ans
}

func pathExistenceQueries1(n int, nums []int, maxDiff int, queries [][]int) []bool {
	idx := []int{}
	for i := range n - 1 {
		if nums[i+1]-nums[i] > maxDiff {
			idx = append(idx, i) // 间断点
		}
	}

	ans := make([]bool, len(queries))
	for i, q := range queries {
		ans[i] = sort.SearchInts(idx, q[0]) == sort.SearchInts(idx, q[1])
	}
	return ans
}

//

type uf struct {
	fa []int
	cc int
}

func newUnionFind(n int) uf {
	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	return uf{fa, n}
}

func (u uf) find(x int) int {
	if u.fa[x] != x {
		u.fa[x] = u.find(u.fa[x])
	}
	return u.fa[x]
}

func (u *uf) merge(from, to int) (isNewMerge bool) {
	x, y := u.find(from), u.find(to)
	if x == y {
		return false
	}
	u.fa[x] = y
	u.cc--
	return true
}

func (u uf) same(x, y int) bool { return u.find(x) == u.find(y) }

func pathExistenceQueries0(n int, nums []int, maxDiff int, queries [][]int) (ans []bool) {
	u := newUnionFind(n)
	for i := 1; i < len(nums); i++ {
		v, w := nums[i-1], nums[i]
		if w-v <= maxDiff {
			u.merge(i-1, i)
		}
	}
	for _, q := range queries {
		ans = append(ans, u.same(q[0], q[1]))
	}
	return
}
