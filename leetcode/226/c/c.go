package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func canEat(a []int, queries [][]int) (ans []bool) {
	n := len(a)
	sum := make([]int, n+1)
	for i, v := range a {
		sum[i+1] = sum[i] + v
	}
	for _, q := range queries {
		tp, d, dc := q[0], q[1], q[2]
		minTP := sort.SearchInts(sum, d+1) - 1
		maxTP := sort.SearchInts(sum, (d+1)*dc) - 1
		ans = append(ans, minTP <= tp && tp <= maxTP)
	}
	return
}
