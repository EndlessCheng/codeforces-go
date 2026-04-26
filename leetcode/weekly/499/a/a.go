package main

import "slices"

// https://space.bilibili.com/206214
func findValidElements(nums []int) (ans []int) {
	// 标记严格大于其右侧元素的元素
	rightValid := make([]bool, len(nums))
	mx := 0
	for i, x := range slices.Backward(nums) {
		rightValid[i] = x > mx
		mx = max(mx, x)
	}

	mx = 0
	for i, x := range nums {
		if x > mx || rightValid[i] {
			ans = append(ans, x)
		}
		mx = max(mx, x)
	}
	return
}
