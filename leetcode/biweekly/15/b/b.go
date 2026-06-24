package main

import (
	"cmp"
	"slices"
)

// github.com/EndlessCheng/codeforces-go
func removeCoveredIntervals(intervals [][]int) (ans int) {
	// 按区间左端点从小到大排序
	// 区间左端点相同时，按区间右端点从大到小排序，这样会先遍历大区间，再遍历被大区间覆盖的小区间
	slices.SortFunc(intervals, func(a, b []int) int {
		return cmp.Or(a[0]-b[0], b[1]-a[1])
	})

	maxRight := 0 // 已遍历区间中的最大右端点
	for _, p := range intervals {
		// 由于区间左端点是从小到大排序的，遍历过的区间的左端点都 <= 当前区间的左端点
		// 如果当前区间右端点 <= maxRight，说明当前区间被另一个区间覆盖，否则没被覆盖
		if p[1] > maxRight {
			maxRight = p[1]
			ans++ // 当前区间没被覆盖
		}
	}
	return
}
