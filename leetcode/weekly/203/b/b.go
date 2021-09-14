package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func maxCoins(a []int) (ans int) {
	for sort.Ints(a); len(a) > 0; a = a[1 : len(a)-2] {
		ans += a[len(a)-2]
	}
	return
}
