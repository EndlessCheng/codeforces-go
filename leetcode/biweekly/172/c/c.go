package main

import (
	"container/heap"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func maximumScore(nums []int, s string) (ans int64) {
	h := hp{}
	for i, x := range slices.Backward(nums) {
		if s[i] == '1' {
			ans += int64(x)
			heap.Push(&h, x)
		} else if h.Len() > 0 && x > h.IntSlice[0] {
			ans += int64(x - h.IntSlice[0])
			h.IntSlice[0] = x
			heap.Fix(&h, 0)
		}
	}
	return
}

type hp struct{ sort.IntSlice }

func (h *hp) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (hp) Pop() (_ any)  { return }
