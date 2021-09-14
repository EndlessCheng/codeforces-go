package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func reductionOperations(a []int) (ans int) {
	sort.Ints(a)
	for i, k, n := 0, 0, len(a); i < n; k++ {
		st := i
		for ; i < n && a[i] == a[st]; i++ {
		}
		ans += (i - st) * k
	}
	return
}
