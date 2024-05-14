package main

import "slices"

// https://space.bilibili.com/206214
func maxDivScore(nums []int, divisors []int) (ans int) {
	slices.SortFunc(nums, func(a, b int) int { return b - a })
	dup := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			dup++
		}
	}
	slices.Sort(divisors)
	maxCnt := -1
	for _, d := range divisors {
		if (maxCnt-dup+1)*d > nums[0] {
			break
		}
		cnt := 0
		for _, x := range nums {
			if x < d {
				break
			}
			if x%d == 0 {
				cnt++
			}
		}
		if cnt > maxCnt {
			maxCnt, ans = cnt, d
		}
	}
	return
}
