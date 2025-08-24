package main

import "math"

// https://space.bilibili.com/206214
func maxValue(nums []int) []int {
	n := len(nums)
	preMax := make([]int, n)
	preMax[0] = nums[0]
	for i := 1; i < n; i++ {
		preMax[i] = max(preMax[i-1], nums[i])
	}

	ans := make([]int, n)
	sufMin := math.MaxInt
	for i := n - 1; i >= 0; i-- {
		if preMax[i] <= sufMin {
			ans[i] = preMax[i]
		} else {
			ans[i] = ans[i+1]
		}
		sufMin = min(sufMin, nums[i])
	}
	return ans
}

func maxValue2(nums []int) []int {
	n := len(nums)
	preMax := make([]int, n)
	preMax[0] = nums[0]
	for i := 1; i < n; i++ {
		preMax[i] = max(preMax[i-1], nums[i])
	}

	sufMin := math.MaxInt
	for i := n - 1; i >= 0; i-- {
		if preMax[i] > sufMin {
			preMax[i] = preMax[i+1]
		}
		sufMin = min(sufMin, nums[i])
	}
	return preMax
}
