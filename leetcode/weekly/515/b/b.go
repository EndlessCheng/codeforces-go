package main

import (
	"slices"
)

// https://space.bilibili.com/206214
func minPenalty(period int, lights []int, arrivalTime []int) (ans int) {
	mx := slices.Max(lights)
	for _, t := range arrivalTime {
		t %= period
		if t >= mx {
			ans = max(ans, period-t)
		}
	}
	return
}
