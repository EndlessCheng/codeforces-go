package main

import "math"

// https://space.bilibili.com/206214
func uniformArray(nums1 []int) bool {
	// 计算最小偶数、最小奇数
	mn := [2]int{math.MaxInt, math.MaxInt}
	for _, x := range nums1 {
		mn[x&1] = min(mn[x&1], x)
	}

	// 只有偶数 or 偶数 >= 最小的偶数 > 最小的奇数
	// 只有奇数的情况蕴含在 mn[0] > mn[1] 中
	return mn[1] == math.MaxInt || mn[0] > mn[1]
}
