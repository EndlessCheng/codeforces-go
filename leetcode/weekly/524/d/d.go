package main

import (
	"slices"
)

// https://space.bilibili.com/206214
func minSelection(availableIntervals, targetIntervals [][]int) (ans int) {
	// 按照左端点从小到大排序
	slices.SortFunc(availableIntervals, func(a, b []int) int { return a[0] - b[0] })
	slices.SortFunc(targetIntervals, func(a, b []int) int { return a[0] - b[0] })

	curEnd := 0 // 已建造的桥的右端点
	i := 0
	for _, target := range targetIntervals {
		tarL, tarR := target[0], target[1]
		for curEnd < tarR {
			// 当前最靠左的未覆盖位置
			uncover := max(tarL, curEnd)

			// 寻找下一座桥的最远点
			nextEnd := curEnd
			for i < len(availableIntervals) && availableIntervals[i][0] <= uncover {
				nextEnd = max(nextEnd, availableIntervals[i][1])
				i++
			}

			// 无法覆盖 uncover，或者无法继续向右延伸
			if nextEnd < uncover || nextEnd == curEnd {
				return -1
			}

			curEnd = nextEnd // 建桥后，最远可以到达 nextEnd
			ans++
		}
	}
	return
}
