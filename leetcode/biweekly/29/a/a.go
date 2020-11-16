package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func average(a []int) (ans float64) {
	sort.Ints(a)
	sum := 0
	for _, v := range a[1 : len(a)-1] {
		sum += v
	}
	return float64(sum) / float64(len(a)-2)
}
