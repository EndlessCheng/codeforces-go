package main

// https://space.bilibili.com/206214
func minDeletions(nums []int) (ans int) {
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			ans++
		}
	}
	return
}
