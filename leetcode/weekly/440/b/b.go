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
	h := hp{make([]int, k)}
	s := 0
	for i, t := range a {
		if i > 0 && t.x == a[i-1].x {
			ans[t.i] = ans[a[i-1].i]
		} else {
			ans[t.i] = int64(s)
		}
		y := t.y
		if i < k {
			s += y
			h.IntSlice[i] = y
			continue
		}
		if i == k {
			heap.Init(&h)
		}
		if y > h.IntSlice[0] {
			s += y - h.IntSlice[0]
			h.IntSlice[0] = y
			heap.Fix(&h, 0)
		}
	}
	return ans
}

type hp struct{ sort.IntSlice }
func (hp) Push(any)     {}
func (hp) Pop() (_ any) { return }
