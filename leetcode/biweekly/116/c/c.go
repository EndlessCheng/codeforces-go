package main

import "math"

// https://space.bilibili.com/206214
func lengthOfLongestSubsequence(nums []int, target int) int {
	f := make([]int, target+1)
	for i := 1; i <= target; i++ {
		f[i] = math.MinInt
	}
	s := 0
	for _, x := range nums {
		s = min(s+x, target)
		for j := s; j >= x; j-- {
			f[j] = max(f[j], f[j-x]+1)
		}
	}
	if f[target] > 0 {
		return f[target]
	}
	return -1
}

func min(a, b int) int { if b < a { return b }; return a }
func max(a, b int) int { if b > a { return b }; return a }
