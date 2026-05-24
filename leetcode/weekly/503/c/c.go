package main

import "math"

// https://space.bilibili.com/206214
func minOperations(nums []int) int {
	n := len(nums)
	cnt := 0
	p := 0
	for i := 1; i < n && cnt < 2; i++ {
		if nums[i-1] > nums[i] {
			cnt++
			p = i
		}
	}

	if cnt == 0 { // 已是递增
		return 0
	}
	ans := math.MaxInt
	if cnt == 1 && nums[0] > nums[n-1] { // 两个递增段
		ans = min(p, n-p+2)
	}

	cnt = 0
	p = 0
	for i := 1; i < n && cnt < 2; i++ {
		if nums[i-1] < nums[i] {
			cnt++
			p = i
		}
	}

	if cnt == 0 { // 已是递减
		return 1
	}
	if cnt == 1 && nums[0] < nums[n-1] { // 两个递减段
		ans = min(ans, p+1, n-p+1)
	}

	if ans == math.MaxInt {
		return -1
	}
	return ans
}
