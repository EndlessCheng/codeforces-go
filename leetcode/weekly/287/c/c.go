package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func maximumCandies(candies []int, k int64) int {
	mx, sum := 0, 0
	for _, c := range candies {
		mx = max(mx, c)
		sum += c
	}
	// 二分最大的不满足要求的 low+1，那么答案就是 low
	return sort.Search(min(mx, sum/int(k)), func(low int) bool {
		low++
		sum := 0
		for _, candy := range candies {
			sum += candy / low
		}
		return sum < int(k)
	})
}
