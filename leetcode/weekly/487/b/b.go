package main

// https://space.bilibili.com/206214
func finalElement(nums []int) int {
	return max(nums[0], nums[len(nums)-1])
}
