package main

// https://space.bilibili.com/206214
func minOperations(nums []int) (ans int64) {
	for i := 1; i < len(nums); i++ {
		ans += int64(max(nums[i-1]-nums[i], 0))
	}
	return
}
