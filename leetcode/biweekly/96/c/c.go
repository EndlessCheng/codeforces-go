package main

import (
	"container/heap"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func maxScore(nums1, nums2 []int, k int) int64 {
	ids := make([]int, len(nums1))
	for i := range ids {
		ids[i] = i
	}
	slices.SortFunc(ids, func(i, j int) int { return nums2[j] - nums2[i] })

	h := hp{make([]int, k)}
	sum := 0
	for i, idx := range ids[:k] {
		sum += nums1[idx]
		h.IntSlice[i] = nums1[idx]
	}
	heap.Init(&h)

	ans := sum * nums2[ids[k-1]]
	for _, i := range ids[k:] {
		x := nums1[i]
		if x > h.IntSlice[0] {
			sum += x - h.replace(x)
			ans = max(ans, sum*nums2[i])
		}
	}
	return int64(ans)
}

type hp struct{ sort.IntSlice }
func (hp) Push(any)            {}
func (hp) Pop() (_ any)        { return }
func (h hp) replace(v int) int { top := h.IntSlice[0]; h.IntSlice[0] = v; heap.Fix(&h, 0); return top }
