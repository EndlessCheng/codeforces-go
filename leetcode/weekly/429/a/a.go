package main

import "slices"

// https://space.bilibili.com/206214
func minimumOperations(nums []int) int {
	seen := map[int]struct{}{}
	for i, x := range slices.Backward(nums) {
		if _, ok := seen[x]; ok {
			return i/3 + 1
		}
		seen[x] = struct{}{}
	}
	return 0
}
