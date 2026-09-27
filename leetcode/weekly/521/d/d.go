package main

import (
	"math"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func maxEarnings(meetings [][]int) int64 {
	// 按结束时间从小到大排序
	slices.SortFunc(meetings, func(a, b []int) int { return a[1] - b[1] })
	end0 := meetings[0][1]

	// preMax[i+1] = [0,i] 中的 f[j] - end[j] 的前缀最大值
	preMax := make([]int, len(meetings)+1)
	preMax[0] = math.MinInt
	ans := 0
	for i, m := range meetings {
		start, end, revenue := m[0], m[1], m[2]

		f := revenue
		if start >= end0 { // 左边有会议
			j := sort.Search(i, func(j int) bool { return meetings[j][1] > start })
			// 为什么是 j 不是 j+1：上面算的是 > start，-1 后得到 <= start，但由于还要 +1，抵消了
			f += preMax[j] + start
		}
		ans = max(ans, f)

		preMax[i+1] = max(preMax[i], f-end)
	}

	return int64(ans)
}
