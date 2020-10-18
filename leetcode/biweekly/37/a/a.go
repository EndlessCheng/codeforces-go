package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func trimMean(a []int) (ans float64) {
	sort.Ints(a)
	n := len(a)
	m := n / 20
	sum := 0
	for _, v := range a[m : n-m] {
		sum += v
	}
	return float64(sum) / float64(n-2*m)
}
