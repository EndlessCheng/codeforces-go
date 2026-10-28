package main

import "slices"

// https://space.bilibili.com/206214
func countServers(n int, logs [][]int, x int, queries []int) []int {
	slices.SortFunc(logs, func(a, b []int) int { return a[1] - b[1] }) // 按照 time 排序
	idx := make([]int, len(queries))
	for i := range idx {
		idx[i] = i
	}
	slices.SortFunc(idx, func(i, j int) int { return queries[i] - queries[j] })
	cnt := make([]int, n+1) // 每个服务器收到了多少个请求
	notReceives := n        // 没有收到请求的服务器数目
	left, right := 0, 0

	for _, qid := range idx {
		q := queries[qid]
		// 进入窗口
		for right < len(logs) && logs[right][1] <= q {
			i := logs[right][0]
			if cnt[i] == 0 {
				notReceives--
			}
			cnt[i]++
			right++
		}

		// 离开窗口
		for left < len(logs) && logs[left][1] < q-x {
			i := logs[left][0]
			cnt[i]--
			if cnt[i] == 0 {
				notReceives++
			}
			left++
		}

		queries[qid] = notReceives
	}

	return queries
}
