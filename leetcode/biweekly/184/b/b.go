package main

import "slices"

// https://space.bilibili.com/206214
func minEnergy(_, brightness int, intervals [][]int) int64 {
	slices.SortFunc(intervals, func(p, q []int) int { return p[0] - q[0] }) // 按照左端点从小到大排序

	// 56. 合并区间（只计算区间长度之和）
	sumLen := 0
	left, right := 0, -1
	for _, p := range intervals {
		if p[0] <= right { // 左端点在合并区间内，可以合并
			right = max(right, p[1]) // 更新合并区间的右端点
		} else { // 不相交，无法合并
			sumLen += right - left + 1
			left, right = p[0], p[1] // 新的合并区间
		}
	}
	sumLen += right - left + 1

	bulbs := (brightness + 2) / 3 // 至少要开启 bulbs 个灯泡
	return int64(bulbs * sumLen)
}
