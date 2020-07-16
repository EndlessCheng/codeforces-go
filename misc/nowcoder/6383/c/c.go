package main

import (
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func solve(n int, _ int, p []int, pairs []*Point) []int {
	fa := make([]int, n+1)
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
	merge := func(from, to int) { fa[find(from)] = find(to) }
	for _, pp := range pairs {
		merge(p[pp.X-1], p[pp.Y-1]) // 合并同组数字
	}

	group := make([][2][]int, n+1)
	for i, v := range p {
		f := find(v)
		group[f][0] = append(group[f][0], v)
		group[f][1] = append(group[f][1], i)
	}

	ans := make([]int, n)
	for _, vp := range group {
		vs, ps := vp[0], vp[1]
		sort.Ints(vs)
		sort.Ints(ps)
		for i, v := range vs {
			ans[ps[i]] = v
		}
	}
	return ans
}
