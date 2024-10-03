package main

import (
	"math"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func minSpeedOnTime(dist []int, hour float64) int {
	n := len(dist)
	h100 := int(math.Round(hour * 100)) // 下面不会用到任何浮点数
	delta := h100 - (n-1)*100
	if delta <= 0 { // 无法到达终点
		return -1
	}

	maxDist := slices.Max(dist)
	if h100 <= n*100 { // 特判
		// 见题解中的公式
		return max(maxDist, (dist[n-1]*100-1)/delta+1)
	}

	sumDist := 0
	for _, d := range dist {
		sumDist += d
	}
	left := (sumDist*100-1)/h100 + 1 // 也可以初始化成 0（简单写法）
	h := h100 / (n * 100)
	right := (maxDist-1)/h + 1 // 也可以初始化成 maxDist（简单写法）
	return left + sort.Search(right-left, func(v int) bool {
		v += left
		t := 0
		for _, d := range dist[:n-1] {
			t += (d-1)/v + 1
		}
		return (t*v+dist[n-1])*100 <= h100*v
	})
}
