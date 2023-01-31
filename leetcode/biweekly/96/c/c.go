package main

import (
	"container/heap"
	"sort"
)

// https://space.bilibili.com/206214
func maxScore(nums1, nums2 []int, k int) int64 {
	type pair struct{ x, y int }
	a := make([]pair, len(nums1))
	sum := 0
	for i, x := range nums1 {
		a[i] = pair{x, nums2[i]}
	}
	sort.Slice(a, func(i, j int) bool { return a[i].y > a[j].y })

	h := hp{nums2[:k]} // 复用内存
	for i, p := range a[:k] {
		sum += p.x
		h.IntSlice[i] = p.x
	}
	ans := sum * a[k-1].y
	heap.Init(&h)
	for _, p := range a[k:] {
		if p.x > h.IntSlice[0] {
			sum += p.x - h.replace(p.x)
			ans = max(ans, sum*p.y)
		}
	}
	return int64(ans)
}

type hp struct{ sort.IntSlice }
func (hp) Pop() (_ interface{}) { return }
func (hp) Push(interface{})     {}
func (h hp) replace(v int) int  { top := h.IntSlice[0]; h.IntSlice[0] = v; heap.Fix(&h, 0); return top }
func max(a, b int) int { if b > a { return b }; return a }
