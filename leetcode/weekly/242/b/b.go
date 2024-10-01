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

	return 1 + sort.Search(maxDist-1, func(v int) bool {
		v++
		t := 0
		for _, d := range dist[:n-1] {
			t += (d-1)/v + 1
		}
		return (t*v+dist[n-1])*100 <= h100*v
	})
}
