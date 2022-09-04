package main

import (
	"container/heap"
	"sort"
)

// https://space.bilibili.com/206214
func mostBooked(n int, meetings [][]int) (ans int) {
	cnt := make([]int, n)
	idle := hp{make([]int, n)}
	for i := 0; i < n; i++ {
		idle.IntSlice[i] = i
	}
	using := hp2{}
	sort.Slice(meetings, func(i, j int) bool { return meetings[i][0] < meetings[j][0] })
	for _, m := range meetings {
		st, end := m[0], m[1]
		for len(using) > 0 && using[0].end <= st { // 维护在 st 时刻空闲的会议室
			heap.Push(&idle, heap.Pop(&using).(pair).i)
		}
		var i int
		if idle.Len() == 0 {
			p := heap.Pop(&using).(pair) // 没有可用的会议室，那么弹出一个最早结束的会议室
			end += p.end - st // 更新当前会议的结束时间
			i = p.i
		} else {
			i = heap.Pop(&idle).(int)
		}
		cnt[i]++
		heap.Push(&using, pair{end, i}) // 使用一个会议室
	}
	for i, c := range cnt {
		if c > cnt[ans] {
			ans = i
		}
	}
	return
}

type hp struct{ sort.IntSlice }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
type pair struct{ end, i int }
type hp2 []pair
func (h hp2) Len() int { return len(h) }
func (h hp2) Less(i, j int) bool {
	a, b := h[i], h[j]
	return a.end < b.end || a.end == b.end && a.i < b.i
}
func (h hp2) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp2) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp2) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
