package main

import "slices"

// https://space.bilibili.com/206214
func semiOrderedPermutation(nums []int) int {
	n := len(nums)
	p := slices.Index(nums, 1)
	q := slices.Index(nums, n)
	if p < q {
		return p + n - 1 - q
	}
	return p + n - 2 - q // 1 向左移动的时候和 n 交换了一次
}
