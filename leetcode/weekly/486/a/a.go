package main

// https://space.bilibili.com/206214
func minimumPrefixLength(nums []int) int {
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i-1] >= nums[i] {
			return i // 移除前缀 [0, i-1]，长度为 i
		}
	}
	return 0
}
