package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func removeCoveredIntervals(ps [][]int) (ans int) {
	sort.Slice(ps, func(i, j int) bool { a, b := ps[i], ps[j]; return a[0] < b[0] || a[0] == b[0] && a[1] > b[1] })
	mx := 0
	for _, p := range ps {
		if p[1] > mx {
			mx = p[1]
			ans++
		}
	}
	return
}
