package main

import "slices"

// https://space.bilibili.com/206214
func countDays(days int, meetings [][]int) int {
	slices.SortFunc(meetings, func(p, q []int) int { return p[0] - q[0] }) // 按照左端点从小到大排序
	start, end := 1, 0 // 当前合并区间的左右端点
	for _, p := range meetings {
		if p[0] > end { // 不相交
			days -= end - start + 1 // 当前合并区间的长度
			start = p[0] // 下一个合并区间的左端点
		}
		end = max(end, p[1])
	}
	days -= end - start + 1 // 最后一个合并区间的长度
	return days
}
