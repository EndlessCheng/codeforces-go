package main

import (
	"slices"
	"sort"
)

// 二分答案

// github.com/EndlessCheng/codeforces-go
func minimizedMaximum2(n int, quantities []int) int {
	left, right := 1, slices.Max(quantities)
	return left + sort.Search(right-left, func(mx int) bool {
		mx += left
		cnt := 0
		for _, q := range quantities {
			cnt += (q-1)/mx + 1
		}
		return cnt <= n
	})
}

func minimizedMaximum(n int, quantities []int) int {
	check := func(mx int) bool {
		cnt := 0
		for _, q := range quantities {
			cnt += (q + mx - 1) / mx
		}
		return cnt <= n
	}

	left, right := 0, slices.Max(quantities)
	for left+1 < right {
		mid := left + (right-left)/2
		if check(mid) {
			right = mid
		} else {
			left = mid
		}
	}
	return right
}
