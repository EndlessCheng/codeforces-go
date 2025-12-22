package main

import (
	"container/heap"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func maximumScore(nums []int, s string) (ans int64) {
	h := hp{}
	for i, x := range nums {
		heap.Push(&h, x)
		if s[i] == '1' {
			ans += int64(heap.Pop(&h).(int))
		}
	}
	return
}

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp) Push(v any)        { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any          { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }

//

func maximumScore2(nums []int, s string) (ans int64) {
	h := hp2{}
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

type hp2 struct{ sort.IntSlice }

func (h *hp2) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (hp2) Pop() (_ any)  { return }
