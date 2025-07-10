package main

import "slices"

// github.com/EndlessCheng/codeforces-go
func maximumUniqueSubarray(nums []int) (ans int) {
	mx := slices.Max(nums)
	has := make([]bool, mx+1)
	s, left := 0, 0
	for _, x := range nums {
		for has[x] {
			has[nums[left]] = false
			s -= nums[left]
			left++
		}
		has[x] = true
		s += x
		ans = max(ans, s)
	}
	return
}
