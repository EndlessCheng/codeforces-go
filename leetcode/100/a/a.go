package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func isMonotonic(a []int) bool {
	return sort.IntsAreSorted(a) || sort.SliceIsSorted(a, func(i, j int) bool { return a[i] > a[j] })
}
