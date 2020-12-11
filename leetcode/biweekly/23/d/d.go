package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func maxSatisfaction(a []int) (ans int) {
	sort.Ints(a)
	for i, s := len(a)-1, 0; i >= 0 && s+a[i] > 0; i-- {
		s += a[i]
		ans += s
	}
	return
}
