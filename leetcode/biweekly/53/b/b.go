package main

import "slices"

// github.com/EndlessCheng/codeforces-go
func minPairSum(nums []int) (ans int) {
	slices.Sort(nums)
	n := len(nums)
	for i, x := range nums[:n/2] {
		ans = max(ans, x+nums[n-1-i])
	}
	return
}
