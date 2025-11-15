package main

import (
	"container/heap"
	"sort"
)

// github.com/EndlessCheng
func isPossible(target []int) bool {
	sum := 0
	for _, x := range target {
		sum += x
	}

	h := hp{target}
	heap.Init(&h)

	// 如果最大值等于 1，说明所有数都等于 1
	for h.IntSlice[0] > 1 {
		x := h.IntSlice[0]
		sum -= x // 减去 x 后，sum 为其余元素之和
		// sum 不能是 0，这意味着 target 只有一个数且这个数大于 1
		// x 减去 sum 后必须是正数
		if sum == 0 || x <= sum {
			return false
		}
		// 把 x 多次减去 sum，直到 x <= sum 为止
		// 也就是计算 x%sum，但如果 x%sum == 0 则调整为 sum
		x = (x-1)%sum + 1
		sum += x
		h.IntSlice[0] = x
		heap.Fix(&h, 0)
	}
	return true
}

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (hp) Push(any)             {} // 没用到，无需实现
func (hp) Pop() (_ any)         { return }
