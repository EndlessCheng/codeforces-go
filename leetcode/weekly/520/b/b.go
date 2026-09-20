package main

import (
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func countIntersectingIntervals(intervals [][]int) int64 {
	n := len(intervals)
	slices.SortFunc(intervals, func(a, b []int) int { return a[1] - b[1] }) // 按照右端点升序排序

	ans := n * (n - 1) / 2
	for _, p := range intervals {
		start := p[0]
		// 设 j 是最小的满足 intervals[j][1] >= start 的下标
		// 那么 [0, j-1] 中的区间右端点都 < start，这有 j 个
		ans -= sort.Search(n, func(j int) bool { return intervals[j][1] >= start })
	}
	return int64(ans)
}

func countIntersectingIntervals2(intervals [][]int) int64 {
	n := len(intervals)
	starts := make([]int, n)
	ends := make([]int, n)
	for i, p := range intervals {
		starts[i] = p[0]
		ends[i] = p[1]
	}

	slices.Sort(starts)
	slices.Sort(ends)

	ans := n * (n - 1) / 2
	// 对于每个左端点 start，右端点 < start 的区间都与之不相交
	j := 0
	for _, start := range starts {
		for j < n && ends[j] < start {
			j++
		}
		// [0, j-1] 的区间与当前区间不相交，这有 j 个
		ans -= j
	}
	return int64(ans)
}
