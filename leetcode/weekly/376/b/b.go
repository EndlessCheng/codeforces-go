package main

import "slices"

// https://space.bilibili.com/206214
func divideArray(nums []int, k int) (ans [][]int) {
	slices.Sort(nums)
	for a := range slices.Chunk(nums, 3) {
		if a[2]-a[0] > k {
			return nil
		}
		ans = append(ans, a)
	}
	return
}
