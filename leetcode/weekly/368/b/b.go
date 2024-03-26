package main

import "math"

// https://space.bilibili.com/206214
func minimumSum(nums []int) int {
	n := len(nums)
	suf := make([]int, n) // 后缀最小值
	suf[n-1] = nums[n-1]
	for i := n - 2; i > 1; i-- {
		suf[i] = min(suf[i+1], nums[i])
	}

	ans := math.MaxInt
	pre := nums[0] // 前缀最小值
	for j := 1; j < n-1; j++ {
		if pre < nums[j] && nums[j] > suf[j+1] { // 山形
			ans = min(ans, pre+nums[j]+suf[j+1]) // 更新答案
		}
		pre = min(pre, nums[j])
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}
