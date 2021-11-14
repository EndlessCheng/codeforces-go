package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func maxTaskAssign(tasks, workers []int, pills, strength int) int {
	sort.Ints(tasks)
	sort.Ints(workers)
	m := len(workers)
	return sort.Search(min(len(tasks), m), func(k int) bool {
		k++
		i, p, validTasks := 0, pills, []int{}
		for _, w := range workers[m-k:] {
			for ; i < k && tasks[i] <= w+strength; i++ {
				validTasks = append(validTasks, tasks[i]) // 保证吃药时去掉的是 validTasks 的最后一个元素
			}
			if len(validTasks) == 0 {
				return true
			}
			if validTasks[0] <= w {
				validTasks = validTasks[1:] // 不吃药时去掉最小的元素
			} else {
				if p == 0 {
					return true
				}
				p--
				validTasks = validTasks[:len(validTasks)-1]
			}
		}
		return false
	})
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
