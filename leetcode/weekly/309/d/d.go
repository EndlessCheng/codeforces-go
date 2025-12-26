package main

import (
	"container/heap"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func mostBooked(n int, meetings [][]int) (ans int) {
	slices.SortFunc(meetings, func(a, b []int) int { return a[0] - b[0] })

	idle := hp{make([]int, n)} // 会议室编号
	for i := range n {
		idle.IntSlice[i] = i
	}
	using := hp2{}        // (结束时间，会议室编号)
	cnt := make([]int, n) // 会议室的开会次数

	for _, m := range meetings {
		start, end := m[0], m[1]

		// 在 start 时刻空出来的会议室
		for len(using) > 0 && using[0].end <= start {
			heap.Push(&idle, heap.Pop(&using).(pair).i)
		}

		var i int
		if idle.Len() > 0 { // 有空闲的会议室
			i = heap.Pop(&idle).(int)
		} else {
			// 弹出一个最早结束的会议室（若有多个同时结束，弹出编号最小的会议室）
			p := heap.Pop(&using).(pair)
			end += p.end - start // 更新当前会议的结束时间  
			i = p.i
		}

		heap.Push(&using, pair{end, i}) // 使用一个会议室
		cnt[i]++
	}

	for i, c := range cnt {
		if c > cnt[ans] {
			ans = i
		}
	}
	return
}

type hp struct{ sort.IntSlice }
func (h *hp) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }

type pair struct{ end, i int }
type hp2 []pair
func (h hp2) Len() int { return len(h) }
func (h hp2) Less(i, j int) bool {
	a, b := h[i], h[j]
	return a.end < b.end || a.end == b.end && a.i < b.i
}
func (h hp2) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *hp2) Push(v any)   { *h = append(*h, v.(pair)) }
func (h *hp2) Pop() any     { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
