package main

import (
	"math"
	"sort"
)

// 排序+滑动窗口

// github.com/EndlessCheng/codeforces-go
func minimumDifference(a []int, k int) int {
	sort.Ints(a)
	ans := math.MaxInt32
	for i := k - 1; i < len(a); i++ {
		ans = min(ans, a[i]-a[i-k+1])
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
