package main

// https://space.bilibili.com/206214
func firstStableIndex(nums []int, k int) int {
	n := len(nums)
	sufMin := make([]int, n) // 后缀最小值
	sufMin[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		sufMin[i] = min(sufMin[i+1], nums[i])
	}

	preMax := 0 // 前缀最大值
	for i, x := range nums {
		preMax = max(preMax, x)
		if preMax-sufMin[i] <= k {
			return i
		}
	}
	return -1
}
