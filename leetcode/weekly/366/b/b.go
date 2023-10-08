package main

import "sort"

// https://space.bilibili.com/206214
func minProcessingTime(processorTime, tasks []int) (ans int) {
	sort.Ints(processorTime)
	sort.Sort(sort.Reverse(sort.IntSlice(tasks)))
	for i, p := range processorTime {
		ans = max(ans, p+tasks[i*4])
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
