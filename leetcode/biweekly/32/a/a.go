package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func findKthPositive(arr []int, k int) int {
	i := sort.Search(len(arr), func(i int) bool {
		return arr[i]-1-i >= k
	})
	return i + k // 推导过程见题解
}
