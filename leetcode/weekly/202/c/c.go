package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func maxDistance(ps []int, m int) int {
	sort.Ints(ps)
	return sort.Search(1e9, func(d int) bool {
		cnt, p := 1, ps[0]
		for {
			i := sort.SearchInts(ps, p+d+1)
			if i == len(ps) {
				return cnt < m
			}
			cnt++
			p = ps[i]
		}
	})
}
