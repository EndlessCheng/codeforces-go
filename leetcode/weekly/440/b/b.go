package main

import (
	"container/heap"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func findMaxSum(nums1, nums2 []int, k int) []int64 {
	n := len(nums1)
	type tuple struct{ x, y, i int }
	a := make([]tuple, n)
	for i, x := range nums1 {
		a[i] = tuple{x, nums2[i], i}
	}
	slices.SortFunc(a, func(p, q tuple) int { return p.x - q.x })

	ans := make([]int64, n)
	h := &hp{}
	s := 0
	for i, t := range a {
		if i > 0 && t.x == a[i-1].x {
			ans[t.i] = ans[a[i-1].i]
		} else {
			ans[t.i] = int64(s)
		}
		s += t.y
		heap.Push(h, t.y)
		if h.Len() > k {
			s -= heap.Pop(h).(int)
		}
	}
	return ans
}

type hp struct{ sort.IntSlice }
func (h *hp) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
