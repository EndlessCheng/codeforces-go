package main

import (
	"container/heap"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func halveArray(nums []int) (ans int) {
	half := 0
	for i := range nums {
		nums[i] <<= 20
		half += nums[i]
	}

	h := hp{nums}
	heap.Init(&h)
	for half /= 2; half > 0; ans++ {
		half -= h.IntSlice[0] / 2
		h.IntSlice[0] /= 2
		heap.Fix(&h, 0)
	}
	return
}

type hp struct{ sort.IntSlice } // 继承 sort.IntSlice 的方法
func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] } // 最大堆
func (hp) Push(any)             {}
func (hp) Pop() (_ any)         { return }
