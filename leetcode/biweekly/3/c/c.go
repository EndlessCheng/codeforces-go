package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func earliestAcq(a [][]int, n int) (ans int) {
	sort.Slice(a, func(i, j int) bool { return a[i][0] < a[j][0] })
	fa := make([]int, n)
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
	merge := func(from, to int) bool {
		x, y := find(from), find(to)
		if x == y {
			return false
		}
		fa[x] = y
		return true
	}
	left := n - 1
	for _, p := range a {
		if merge(p[1], p[2]) {
			left--
			if left == 0 {
				return p[0]
			}
		}
	}
	return -1
}
