package main

import (
	"container/heap"
	"sort"
)

// https://space.bilibili.com/206214
func maxKelements(nums []int, k int) (ans int64) {
	h := hp{nums}
	heap.Init(&h) // 原地堆化
	for ; k > 0; k-- {
		ans += int64(h.IntSlice[0])
		h.IntSlice[0] = (h.IntSlice[0] + 2) / 3
		heap.Fix(&h, 0)
	}
	return
}

type hp struct{ sort.IntSlice }
func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] } // 最大堆
func (hp) Push(any)             {}
func (hp) Pop() (_ any)         { return }
