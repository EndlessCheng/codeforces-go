package main

import (
	"container/heap"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool    { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp) Push(interface{})     {}
func (h *hp) Pop() (_ interface{}) { return }

func minimumDeviation(a []int) (ans int) {
	mi := int(1e9)
	for i, v := range a {
		a[i] <<= v & 1
		mi = min(mi, a[i])
	}
	h := &hp{a}
	heap.Init(h)
	ans = 1e9
	for {
		mx := a[0]
		ans = min(ans, mx-mi)
		if mx&1 > 0 {
			return
		}
		a[0] >>= 1
		mi = min(mi, a[0])
		heap.Fix(h, 0)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
