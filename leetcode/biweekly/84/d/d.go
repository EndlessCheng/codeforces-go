package main

// https://space.bilibili.com/206214
func minimumReplacement(nums []int) (ans int64) {
	m := nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		k := (nums[i] - 1) / m
		ans += int64(k)
		m = nums[i] / (k + 1)
	}
	return
}
