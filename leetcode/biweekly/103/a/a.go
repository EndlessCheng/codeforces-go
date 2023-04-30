package main

// https://space.bilibili.com/206214
func maximizeSum(nums []int, k int) int {
	max := nums[0]
	for _, x := range nums[1:] {
		if x > max {
			max = x
		}
	}
	return (max*2 + k - 1) * k / 2
}
