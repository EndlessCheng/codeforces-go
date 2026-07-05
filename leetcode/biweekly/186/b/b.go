package main

// https://space.bilibili.com/206214
func maxValidPairSum(nums []int, k int) (ans int) {
	mx := 0
	for j := k; j < len(nums); j++ {
		mx = max(mx, nums[j-k]) // nums[i] 的最大值
		ans = max(ans, mx+nums[j])
	}
	return
}
