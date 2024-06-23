package main

// https://space.bilibili.com/206214
func minOperations(nums []int) (ans int) {
	n := len(nums)
	for i, x := range nums[:n-2] {
		if x == 0 {
			nums[i+1] ^= 1
			nums[i+2] ^= 1
			ans++
		}
	}
	if nums[n-2] == 0 || nums[n-1] == 0 {
		return -1
	}
	return
}
