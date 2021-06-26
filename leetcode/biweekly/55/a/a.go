package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func canBeIncreasing(a []int) bool {
	if strictlyIncreasing(a) {
		return true
	}
	for i := range a {
		b := append([]int(nil), a...)
		b = append(b[:i], b[i+1:]...)
		if strictlyIncreasing(b) {
			return true
		}
	}
	return false
}

func strictlyIncreasing(a []int) bool {
	return sort.SliceIsSorted(a, func(i, j int) bool { return a[i] <= a[j] })
}
