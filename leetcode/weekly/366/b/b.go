package main

import "slices"

// https://space.bilibili.com/206214
func minProcessingTime(processorTime, tasks []int) (ans int) {
	slices.Sort(processorTime)
	slices.SortFunc(tasks, func(a, b int) int { return b - a })
	for i, p := range processorTime {
		ans = max(ans, p+tasks[i*4])
	}
	return
}
