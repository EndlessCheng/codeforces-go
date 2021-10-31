package main

import (
	"container/heap"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func maxTwoEvents(events [][]int) (ans int) {
	sort.Slice(events, func(i, j int) bool { return events[i][0] < events[j][0] }) // 按开始时间排序
	h := hp{}
	maxVal := 0 // 另一个活动的最大价值
	for _, e := range events {
		start, end, val := e[0], e[1], e[2]
		for len(h) > 0 && h[0].end < start { // 如果结束时间早于当前活动开始时间
			maxVal = max(maxVal, heap.Pop(&h).(pair).val) // 更新另一个活动的最大价值
		}
		ans = max(ans, maxVal+val) // 至多参加两个活动
		heap.Push(&h, pair{end, val})
	}
	return
}

type pair struct{ end, val int }
type hp []pair
func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].end < h[j].end }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
