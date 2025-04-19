package main

import (
	"slices"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func maxTaskAssign(tasks, workers []int, pills, strength int) int {
	slices.Sort(tasks)
	slices.Sort(workers)
	m := len(workers)
	ans := sort.Search(min(len(tasks), m), func(k int) bool {
		k++
		// 贪心：用最强的 k 名工人，完成最简单的 k 个任务
		i, p := 0, pills
		validTasks := []int{}
		for _, w := range workers[m-k:] { // 枚举工人
			// 在吃药的情况下，把能完成的任务记录到 validTasks 中
			for ; i < k && tasks[i] <= w+strength; i++ {
				validTasks = append(validTasks, tasks[i])
			}
			// 即使吃药也无法完成任务
			if len(validTasks) == 0 {
				return true
			}
			// 无需吃药就能完成（最简单的）任务
			if w >= validTasks[0] {
				validTasks = validTasks[1:]
				continue
			}
			// 必须吃药
			if p == 0 { // 没药了
				return true
			}
			p--
			// 完成（能完成的）最难的任务
			validTasks = validTasks[:len(validTasks)-1]
		}
		return false
	})
	return ans
}
