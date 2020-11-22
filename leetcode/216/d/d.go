package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func minimumEffort(t [][]int) (ans int) {
	sort.Slice(t, func(i, j int) bool { a, b := t[i], t[j]; return a[0]-a[1] < b[0]-b[1] })
	ans = sort.Search(1e9+10, func(x int) bool {
		for _, p := range t {
			if x < p[1] {
				return false
			}
			x -= p[0]
		}
		return true
	})
	return
}
