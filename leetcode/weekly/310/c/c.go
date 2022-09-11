package main

import (
	"container/heap"
	"sort"
)

// https://space.bilibili.com/206214
func minGroups(intervals [][]int) int {
	h := hp{}
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	for _, p := range intervals {
		if h.Len() == 0 || p[0] <= h.IntSlice[0] {
			heap.Push(&h, p[1])
		} else {
			h.IntSlice[0] = p[1]
			heap.Fix(&h, 0)
		}
	}
	return h.Len()
}

type hp struct{ sort.IntSlice }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
