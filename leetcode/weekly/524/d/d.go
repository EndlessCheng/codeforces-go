package main

import "slices"

// https://space.bilibili.com/206214
// 56. 合并区间
func merge(intervals [][]int) (ans [][]int) {
	slices.SortFunc(intervals, func(p, q []int) int { return p[0] - q[0] }) // 按照左端点从小到大排序

	for _, p := range intervals {
		m := len(ans)
		if m > 0 && p[0] <= ans[m-1][1] { // 左端点在合并区间内，可以合并
			ans[m-1][1] = max(ans[m-1][1], p[1]) // 更新合并区间的右端点
		} else { // 不相交，无法合并
			ans = append(ans, p) // 新的合并区间
		}
	}
	return
}

func minSelection(availableIntervals, targetIntervals [][]int) (ans int) {
	slices.SortFunc(availableIntervals, func(a, b []int) int { return a[0] - b[0] })

	curEnd := 0  // 已建造的桥的右端点
	nextEnd := 0 // 下一座桥的右端点的最大值
	i := 0
	for _, target := range merge(targetIntervals) {
		tarL, tarR := target[0], target[1]
		for curEnd < tarR {
			// 当前最靠左的未覆盖位置
			uncover := max(tarL, curEnd)

			// 寻找下一座桥的最远点
			for i < len(availableIntervals) && availableIntervals[i][0] <= uncover {
				nextEnd = max(nextEnd, availableIntervals[i][1])
				i++
			}

			// 无法覆盖 uncover，或者无法继续向右延伸
			if nextEnd < uncover || nextEnd == uncover && uncover < tarR {
				return -1
			}

			curEnd = nextEnd // 建桥后，最远可以到达 nextEnd
			ans++
		}
	}
	return
}
