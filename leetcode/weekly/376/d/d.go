package main

import (
	"slices"
)

// https://space.bilibili.com/206214
func maxFrequencyScore(nums []int, k int64) int {
	slices.Sort(nums)
	ans, left := 0, 0
	s := int64(0)
	for right, x := range nums {
		s += int64(x - nums[(left+right)/2])
		for s > k {
			s += int64(nums[left] - nums[(left+right+1)/2])
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

func maxFrequencyScore2(nums []int, K int64) (ans int) {
	k := int(K)
	slices.Sort(nums)

	n := len(nums)
	sum := make([]int, n+1)
	for i, v := range nums {
		sum[i+1] = sum[i] + v
	}

	// 把 nums[l] 到 nums[r] 都变成 nums[i]
	distanceSum := func(l, i, r int) int {
		left := nums[i]*(i-l) - (sum[i] - sum[l])
		right := sum[r+1] - sum[i+1] - nums[i]*(r-i)
		return left + right
	}

	left := 0
	for i := range nums {
		for distanceSum(left, (left+i)/2, i) > k {
			left++
		}
		ans = max(ans, i-left+1)
	}
	return
}
