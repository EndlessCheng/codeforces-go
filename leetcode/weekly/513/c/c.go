package main

import "sort"

// https://space.bilibili.com/206214
func countTasks(tasks, shifts []int) []int {
	n := len(tasks)
	// 原地计算 tasks 的前缀和
	for i := 1; i < n; i++ {
		tasks[i] += tasks[i-1]
	}

	t := 0
	for i, shift := range shifts {
		t += shift
		if t >= tasks[n-1] { // 完成所有任务
			t = 0
			shifts[i] = 0
		} else {
			shifts[i] = n - sort.SearchInts(tasks, t+1)
		}
	}
	return shifts
}
