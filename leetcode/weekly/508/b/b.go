package main

import "slices"

// https://space.bilibili.com/206214
func filterOccupiedIntervals(occupiedIntervals [][]int, freeStart int, freeEnd int) (ans [][]int) {
	slices.SortFunc(occupiedIntervals, func(a, b []int) int { return a[0] - b[0] }) // 按照左端点从小到大排序

	add := func(l, r int) {
		if r < freeStart || l > freeEnd { // 不相交
			ans = append(ans, []int{l, r})
			return
		}
		if l < freeStart {
			ans = append(ans, []int{l, freeStart - 1}) // 余留前缀
		}
		if r > freeEnd {
			ans = append(ans, []int{freeEnd + 1, r}) // 余留后缀
		}
	}

	left := occupiedIntervals[0][0]
	maxR := occupiedIntervals[0][1]
	for _, p := range occupiedIntervals[1:] { // 从第二个区间开始
		l, r := p[0], p[1]
		if l-1 > maxR { // 发现一个新区间
			add(left, maxR) // 先把旧的加入答案
			left = l        // 记录新区间左端点
		}
		maxR = max(maxR, r)
	}
	add(left, maxR)

	return
}
