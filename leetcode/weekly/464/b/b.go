package main

import "slices"

// https://space.bilibili.com/206214
func partitionArray1(nums []int, k int) bool {
	n := len(nums)
	if n%k > 0 {
		return false
	}
	cnt := map[int]int{}
	mx := 0
	for _, x := range nums {
		cnt[x]++
		mx = max(mx, cnt[x])
	}
	return mx <= n/k
}

func partitionArray(nums []int, k int) bool {
	n := len(nums)
	if n%k > 0 {
		return false
	}
	cnt := make([]int, slices.Max(nums)+1)
	for _, x := range nums {
		cnt[x]++
		if cnt[x] > n/k {
			return false
		}
	}
	return true
}
