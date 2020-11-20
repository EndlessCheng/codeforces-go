package main

import "sort"

// github.com/EndlessCheng/codeforces-go
type uf struct {
	fa []int
	sz []int
}

func newUnionFind(n int) uf {
	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	sz := make([]int, n)
	for i := range sz {
		sz[i] = 1
	}
	return uf{fa, sz}
}

func (u uf) find(x int) int {
	if u.fa[x] != x {
		u.fa[x] = u.find(u.fa[x])
	}
	return u.fa[x]
}

func (u uf) merge(from, to int) (mergedSize int) {
	x, y := u.find(from), u.find(to)
	u.fa[x] = y
	u.sz[y] += u.sz[x]
	return u.sz[y] - 1
}

func getMinimums(a []int) []int {
	pos := map[int][]int{}
	for i, v := range a {
		pos[v] = append(pos[v], i)
	}
	keys := make([]int, 0, len(pos))
	for k := range pos {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	n := len(a)
	uf := newUnionFind(n + 1)
	ans := make([]int, n+1)
	cur := 1
	for _, v := range keys {
		maxSize := 0
		for _, p := range pos[v] {
			if sz := uf.merge(p, p+1); sz > maxSize {
				maxSize = sz
			}
		}
		for ; cur <= maxSize; cur++ {
			ans[cur] = v
		}
	}
	return ans[1:]
}
