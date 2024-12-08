package main

// https://space.bilibili.com/206214
func constructTransformedArray(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	for i, x := range nums {
		result[i] = nums[((i+x)%n+n)%n]
	}
	return result
}
