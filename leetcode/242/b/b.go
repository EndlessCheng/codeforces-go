package main

import (
	"math"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func minSpeedOnTime(dist []int, hour float64) int {
	h100 := int(math.Round(hour * 100))
	n := len(dist)
	if h100 <= (n-1)*100 { // hour 必须严格大于 n-1
		return -1
	}
	return 1 + sort.Search(1e7-1, func(v int) bool {
		v++
		h := n - 1
		for _, d := range dist[:n-1] {
			h += (d - 1) / v
		}
		return (h*v+dist[n-1])*100 <= h100*v
	})
}
