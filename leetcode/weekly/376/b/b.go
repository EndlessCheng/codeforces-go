package main

import "slices"

// https://space.bilibili.com/206214
func divideArray(nums []int, k int) (ans [][]int) {
	slices.Sort(nums)
	for i := 2; i < len(nums); i += 3 {
		if nums[i]-nums[i-2] > k {
			return nil
		}
		ans = append(ans, nums[i-2:i+1])
	}
	return
}
