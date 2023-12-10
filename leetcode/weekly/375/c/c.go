package main

import "slices"

// https://space.bilibili.com/206214
func countSubarrays(nums []int, k int) (ans int64) {
	mx := slices.Max(nums)
	cntMx, left := 0, 0
	for _, x := range nums {
		if x == mx {
			cntMx++
		}
		for cntMx == k {
			if nums[left] == mx {
				cntMx--
			}
			left++
		}
		ans += int64(left)
	}
	return
}
