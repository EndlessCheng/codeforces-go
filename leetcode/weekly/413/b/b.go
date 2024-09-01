package main

import (
	"container/heap"
	"sort"
)

// https://space.bilibili.com/206214
func resultsArray(queries [][]int, k int) []int {
	m := len(queries)
	ans := make([]int, m)
	if m < k {
		for i := range ans {
			ans[i] = -1
		}
		return ans
	}
	h := hp{make([]int, k)}
	for i, q := range queries[:k] {
		h.IntSlice[i] = abs(q[0]) + abs(q[1])
		ans[i] = -1
	}
	heap.Init(&h)
	ans[k-1] = h.IntSlice[0]
	for i := k; i < m; i++ {
		q := queries[i]
		d := abs(q[0]) + abs(q[1])
		if d < h.IntSlice[0] {
			h.IntSlice[0] = d
			heap.Fix(&h, 0)
		}
		ans[i] = h.IntSlice[0]
	}
	return ans
}

type hp struct{ sort.IntSlice }
func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (hp) Push(any)             {}
func (hp) Pop() (_ any)         { return }
func abs(x int) int { if x < 0 { return -x }; return x }