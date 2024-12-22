package main

import "slices"

// https://space.bilibili.com/206214
type pair struct{ l, r int }

func check(intervals []pair) bool {
	// 按照左端点从小到大排序
	slices.SortFunc(intervals, func(a, b pair) int { return a.l - b.l })
	cnt, maxR := 0, 0
	for _, p := range intervals {
		if p.l >= maxR { // 新区间
			cnt++
		}
		maxR = max(maxR, p.r) // 更新右端点最大值
	}
	return cnt >= 3 // 也可以在循环中提前退出
}

func checkValidCuts(_ int, rectangles [][]int) bool {
	a := make([]pair, len(rectangles))
	b := make([]pair, len(rectangles))
	for i, rect := range rectangles {
		a[i] = pair{rect[0], rect[2]}
		b[i] = pair{rect[1], rect[3]}
	}
	return check(a) || check(b)
}
