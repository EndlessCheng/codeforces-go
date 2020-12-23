package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func minCostToSupplyWater(n int, wells []int, edges [][]int) (ans int) {
	for i, v := range wells {
		edges = append(edges, []int{0, i + 1, v})
	}
	sort.Slice(edges, func(i, j int) bool { return edges[i][2] < edges[j][2] })

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
	for _, e := range edges {
		if fv, fw := find(e[0]), find(e[1]); fv != fw {
			ans += e[2]
			fa[fv] = fw
		}
	}
	return
}
