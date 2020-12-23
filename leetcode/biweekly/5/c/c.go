package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func minimumCost(n int, edges [][]int) (ans int) {
	sort.Slice(edges, func(i, j int) bool { return edges[i][2] < edges[j][2] })
	var fa []int
	initFa := func(n int) {
		fa = make([]int, n)
		for i := range fa {
			fa[i] = i
		}
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	initFa(n)
	cntE := 0
	for _, e := range edges {
		if fv, fw := find(e[0]-1), find(e[1]-1); fv != fw {
			ans += e[2]
			fa[fv] = fw
			cntE++
		}
	}
	if cntE < n-1 {
		return -1
	}
	return
}
