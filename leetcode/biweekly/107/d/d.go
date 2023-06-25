package main

import "sort"

// https://space.bilibili.com/206214
func countServers(n int, logs [][]int, x int, queries []int) []int {
	type pair struct{ q, i int }
	qs := make([]pair, len(queries))
	for i, q := range queries {
		qs[i] = pair{q, i}
	}
	sort.Slice(qs, func(i, j int) bool { return qs[i].q < qs[j].q })
	sort.Slice(logs, func(i, j int) bool { return logs[i][1] < logs[j][1] }) // 按照 time 排序

	ans := make([]int, len(queries))
	cnt := make([]int, n+1)
	outOfRange, left, right := n, 0, 0
	for _, p := range qs {
		for ; right < len(logs) && logs[right][1] <= p.q; right++ { // 进入窗口
			i := logs[right][0]
			if cnt[i] == 0 {
				outOfRange--
			}
			cnt[i]++
		}
		for ; left < len(logs) && logs[left][1] < p.q-x; left++ { // 离开窗口
			i := logs[left][0]
			cnt[i]--
			if cnt[i] == 0 {
				outOfRange++
			}
		}
		ans[p.i] = outOfRange
	}
	return ans
}
