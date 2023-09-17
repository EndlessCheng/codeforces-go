package main

import "sort"

// https://space.bilibili.com/206214
func minLengthAfterRemovals(nums []int) int {
	n := len(nums)
	x := nums[n/2]
	maxCnt := sort.SearchInts(nums, x+1) - sort.SearchInts(nums, x)
	return max(maxCnt*2-n, n%2)
}

func max(a, b int) int { if b > a { return b }; return a }

func minLengthAfterRemovals2(nums []int) int {
	n := len(nums)
	maxCnt, cnt := 0, 0
	for i := 0; i < n; i++ {
		cnt++
		if i == n-1 || nums[i] != nums[i+1] {
			maxCnt = max(maxCnt, cnt)
			cnt = 0
		}
	}
	if maxCnt*2 > n {
		return maxCnt*2 - n
	}
	return n % 2
}
