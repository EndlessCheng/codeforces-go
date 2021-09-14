package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func maxProductDifference(a []int) int {
	sort.Ints(a)
	n := len(a)
	return a[n-1]*a[n-2] - a[0]*a[1]
}
