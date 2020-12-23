package main

import (
	"container/heap"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
type hp struct{ sort.IntSlice }

func (h *hp) Push(interface{})     {}
func (h *hp) Pop() (_ interface{}) { h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]; return }

func minBuildTime(a []int, split int) (ans int) {
	h := &hp{a}
	heap.Init(h)
	for h.Len() > 1 {
		heap.Pop(h)
		h.IntSlice[0] += split
		heap.Fix(h, 0)
	}
	return h.IntSlice[0]
}
