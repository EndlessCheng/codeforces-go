package main

import "sort"

// https://space.bilibili.com/206214/dynamic
func partitionArray(nums []int, k int) int {
	sort.Ints(nums)
	ans, min := 1, nums[0]
	for _, num := range nums {
		if num-min > k {
			ans++
			min = num
		}
	}
	return ans
}
