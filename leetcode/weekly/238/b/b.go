package main

import "slices"

// github.com/EndlessCheng/codeforces-go
func maxFrequency(nums []int, k int) (ans int) {
	slices.Sort(nums)
	sum, left := 0, 0
	for right, x := range nums {
		sum += x
		for (right-left+1)*x-sum > k {
			sum -= nums[left]
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}
