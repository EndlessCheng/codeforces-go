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

func isPossible(a []int) bool {
	sum := 0
	for _, v := range a {
		sum += v
	}
	h := hp{a}
	heap.Init(&h)
	for {
		max := h.IntSlice[0]
		sum -= max
		if max == 1 || sum == 1 {
			return true
		}
		if max < sum || sum == 0 || max%sum == 0 {
			return false
		}
		max %= sum
		sum += max
		h.IntSlice[0] = max
		heap.Fix(&h, 0)
	}
}
