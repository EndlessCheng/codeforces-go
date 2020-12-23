package main

import (
	"container/heap"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
type hp struct{ sort.IntSlice }

func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
func (h *hp) push(v int)         { heap.Push(h, v) }
func (h *hp) pop() int           { return heap.Pop(h).(int) }

func connectSticks(sticks []int) (ans int) {
	h := hp{sticks}
	heap.Init(&h)
	for h.Len() > 1 {
		cost := h.pop() + h.pop()
		ans += cost
		h.push(cost)
	}
	return
}
