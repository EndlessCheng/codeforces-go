package main

import "math"

// https://space.bilibili.com/206214
func minOperations(nums1, nums2 []int) int64 {
	target := nums2[len(nums2)-1]
	ans := 1 // 把元素追加到 nums1 的末尾需要一次操作
	mn := math.MaxInt
	for i, x := range nums1 {
		y := nums2[i]
		if x > y {
			x, y = y, x
		}
		ans += y - x
		// 如果 target 在 [x,y] 中，那么在从 x 变成 y 的过程中，可以顺带把 target 追加到 nums1 的末尾，代价为 0
		// 如果 target < x，代价为 x-target
		// 如果 target > y，代价为 target-y
		mn = min(mn, max(x-target, target-y))
	}
	return int64(ans + max(mn, 0))
}
