package main

// https://space.bilibili.com/206214
func minOperations(nums []int, k int) int {
	all := 2<<k - 2 // 1~k
	set := 0
	for i := len(nums) - 1; ; i-- {
		set |= 1 << nums[i]
		if set&all == all {
			return len(nums) - i
		}
	}
}
