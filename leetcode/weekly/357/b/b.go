package main

// https://space.bilibili.com/206214
func canSplitArray(nums []int, m int) bool {
	n := len(nums)
	if n <= 2 {
		return true
	}
	for i := 1; i < n; i++ {
		if nums[i-1]+nums[i] >= m {
			return true
		}
	}
	return false
}
