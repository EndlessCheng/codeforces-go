package main

import "sort"

func maximumProduct(a []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n := len(a)
	sort.Ints(a)
	return max(a[n-1]*a[n-2]*a[n-3], a[0]*a[1]*a[n-1])
}
