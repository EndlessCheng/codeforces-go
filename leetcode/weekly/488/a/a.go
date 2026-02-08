package main

// https://space.bilibili.com/206214
func dominantIndices(nums []int) (ans int) {
	n := len(nums)
	sufSum := 0
	for i := n - 2; i >= 0; i-- {
		sufSum += nums[i+1]
		if nums[i]*(n-1-i) > sufSum {
			ans++
		}
	}
	return
}
