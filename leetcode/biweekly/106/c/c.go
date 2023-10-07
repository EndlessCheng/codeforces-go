package main

import "sort"

// https://space.bilibili.com/206214
func sumDistance(nums []int, s string, d int) (ans int) {
	const mod = 1_000_000_007
	for i, c := range s {
		nums[i] += d * int(c&2-1) // L=-1, R=1
	}
	sort.Ints(nums)
	sum := 0
	for i, x := range nums {
		ans = (ans + i*x - sum) % mod
		sum += x
	}
	return
}
