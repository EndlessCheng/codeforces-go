package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func maximumElementAfterDecrementingAndRearranging(a []int) int {
	sort.Ints(a)
	a[0] = 1
	n := len(a)
	for i := 1; i < n; i++ {
		if a[i]-a[i-1] > 1 {
			a[i] = a[i-1] + 1
		}
	}
	return a[n-1]
}
