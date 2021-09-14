package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func isMonotonic(a []int) bool {
	return sort.IntsAreSorted(a) || sort.IsSorted(sort.Reverse(sort.IntSlice(a)))
}
