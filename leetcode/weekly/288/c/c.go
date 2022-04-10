package main

import (
	"container/heap"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func maximumProduct(nums []int, k int) int {
	h := hp{nums}
	for heap.Init(&h); k > 0; k-- {
		h.IntSlice[0]++ // 每次给最小的加一
		heap.Fix(&h, 0)
	}
	ans := 1
	for _, num := range nums {
		ans = ans * num % (1e9 + 7)
	}
	return ans
}

type hp struct{ sort.IntSlice }
func (hp) Push(interface{})     {}
func (hp) Pop() (_ interface{}) { return }
