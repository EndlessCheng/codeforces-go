package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func maximumCandies(candies []int, k int64) int {
	return sort.Search(1e7, func(size int) bool {
		size++
		cnt := int64(0)
		for _, candy := range candies {
			cnt += int64(candy / size)
		}
		return cnt < k
	})
}
