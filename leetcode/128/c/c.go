package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func shipWithinDays(a []int, d int) int {
	return sort.Search(25e6, func(sz int) bool {
		for i, n, c := 0, len(a), 0; i < n; {
			for s := 0; i < n && s+a[i] <= sz; i++ {
				s += a[i]
			}
			if c++; c > d {
				return false
			}
		}
		return true
	})
}
