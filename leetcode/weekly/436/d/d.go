package main

import (
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func maxScore(points []int, m int) int64 {
	u1 := (m + 1) / 2 * slices.Min(points)
	u2 := m / len(points) * slices.Max(points)
	ans := sort.Search(min(u1, u2), func(low int) bool {
		// 二分最小的不满足要求的 low+1，即可得到最大的满足要求的 low
		low++
		left := m
		pre := 0
		for i, p := range points {
			k := (low-1)/p + 1 - pre          // 还需要操作的次数
			if i == len(points)-1 && k <= 0 { // 最后一个数已经满足要求
				break
			}
			k = max(k, 1)   // 至少要走 1 步
			left -= k*2 - 1 // 左右横跳
			if left < 0 {
				return true
			}
			pre = k - 1 // 右边那个数已经操作 k-1 次
		}
		return false
	})
	return int64(ans)
}
