package main

import "slices"

// https://space.bilibili.com/206214
func minimumOperations(nums []int) int {
	seen := map[int]bool{}
	for i, x := range slices.Backward(nums) {
		if seen[x] {
			return i/3 + 1
		}
		seen[x] = true
	}
	return 0
}
