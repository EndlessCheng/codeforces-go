package main

import (
	"math"
	"slices"
)

// https://space.bilibili.com/206214
func filterOccupiedIntervals(occupiedIntervals [][]int, freeStart int, freeEnd int) (ans [][]int) {
	slices.SortFunc(occupiedIntervals, func(a, b []int) int { return a[0] - b[0] }) // 按照左端点从小到大排序

	left, right := math.MaxInt, 0
	for i, p := range occupiedIntervals {
		left = min(left, p[0])
		right = max(right, p[1])
		if i == len(occupiedIntervals)-1 || occupiedIntervals[i+1][0]-1 > right {
			if right < freeStart || left > freeEnd { // 不相交
				ans = append(ans, []int{left, right})
			} else {
				if left < freeStart {
					ans = append(ans, []int{left, freeStart - 1}) // 余留前缀
				}
				if right > freeEnd {
					ans = append(ans, []int{freeEnd + 1, right}) // 余留后缀
				}
			}
			left = math.MaxInt
		}
	}

	return
}
