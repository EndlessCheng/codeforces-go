package main

import "sort"

func maxEvents(a [][]int) (ans int) {
	fa := [1e5 + 1]int{}
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
	sort.Slice(a, func(i, j int) bool { return a[i][1] < a[j][1] })
	for _, e := range a {
		if faL := find(e[0]); faL <= e[1] {
			ans++
			fa[faL] = faL + 1
		}
	}
	return
}
