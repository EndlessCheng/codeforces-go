package main

import "sort"

// 在线的话可以用最小生成树+树链剖分

// github.com/EndlessCheng/codeforces-go
type uf struct {
	fa []int
}

func newUnionFind(n int) uf {
	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	return uf{fa}
}

func (u uf) find(x int) int {
	if u.fa[x] != x {
		u.fa[x] = u.find(u.fa[x])
	}
	return u.fa[x]
}
func (u uf) merge(from, to int) { u.fa[u.find(from)] = u.find(to) }
func (u uf) same(x, y int) bool { return u.find(x) == u.find(y) }

func distanceLimitedPathsExist(n int, es, qs [][]int) (ans []bool) {
	sort.Slice(es, func(i, j int) bool { return es[i][2] < es[j][2] })
	ans = make([]bool, len(qs))
	u := newUnionFind(n)
	for i := range qs {
		qs[i] = append(qs[i], i)
	}
	sort.Slice(qs, func(i, j int) bool { return qs[i][2] < qs[j][2] })
	for _, q := range qs {
		for len(es) > 0 && es[0][2] < q[2] {
			e := es[0]
			es = es[1:]
			u.merge(e[0], e[1])
		}
		ans[q[3]] = u.same(q[0], q[1])
	}
	return
}
