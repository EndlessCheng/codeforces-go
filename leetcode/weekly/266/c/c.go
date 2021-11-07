package main

import "sort"

// 二分答案

// github.com/EndlessCheng/codeforces-go
func minimizedMaximum(n int, quantities []int) int {
	return sort.Search(1e5, func(max int) bool {
		cnt := 0
		for _, q := range quantities {
			cnt += (q + max) / (max + 1)
		}
		return cnt <= n
	}) + 1
}
