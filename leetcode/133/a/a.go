package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func twoCitySchedCost(a [][]int) (ans int) {
	for _, p := range a {
		ans += p[0]
		p[1] -= p[0]
	}
	sort.Slice(a, func(i, j int) bool { return a[i][1] < a[j][1] })
	for _, p := range a[:len(a)/2] {
		ans += p[1]
	}
	return
}
