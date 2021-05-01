package main

import (
	"container/heap"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
type hp struct{ sort.IntSlice }

func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }

type SeatManager struct{}

var h hp

func Constructor(n int) (_ SeatManager) {
	h = hp{make([]int, n)}
	for i := range h.IntSlice {
		h.IntSlice[i] = i + 1
	}
	heap.Init(&h)
	return
}
func (SeatManager) Reserve() int    { return heap.Pop(&h).(int) }
func (SeatManager) Unreserve(x int) { heap.Push(&h, x) }
