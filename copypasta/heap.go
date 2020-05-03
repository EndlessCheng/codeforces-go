package copypasta

import (
	"container/heap"
	"sort"
)

// 下面这些都是最小堆

type hp struct{ sort.IntSlice }

//func (h hp) Less(i, j int) bool    { return h.IntSlice[i] > h.IntSlice[j] } // 最大堆
func (h *hp) Push(v interface{})   { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() (v interface{}) { a := h.IntSlice; h.IntSlice, v = a[:len(a)-1], a[len(a)-1]; return }
func (h *hp) push(v int)           { heap.Push(h, v) }
func (h *hp) pop() int             { return heap.Pop(h).(int) }
func (h hp) empty() bool           { return len(h.IntSlice) == 0 }
func (h hp) top() int              { return h.IntSlice[0] }

//

type hp64 []int64 // 自定义类型

func (h hp64) Len() int              { return len(h) }
func (h hp64) Less(i, j int) bool    { return h[i] < h[j] } // > 为最大堆
func (h hp64) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp64) Push(v interface{})   { *h = append(*h, v.(int64)) }
func (h *hp64) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
func (h *hp64) push(v int64)         { heap.Push(h, v) }
func (h *hp64) pop() int64           { return heap.Pop(h).(int64) }
func (h hp64) empty() bool           { return len(h) == 0 }
func (h hp64) top() int64            { return h[0] }

//

func heapCollections() {
	// 对顶堆求动态中位数（一个数组前 2k+1 项的中位数）
	// https://www.luogu.com.cn/problem/P1168
	dynamicMedians := func(a []int) []int {
		n := len(a)
		medians := make([]int, 1, (n+1)/2)
		medians[0] = a[0]
		small, big := &hp{}, &hp{}
		big.push(a[0]) // 下面保证 big.size() == small.size() || big.size()-1 == small.size()
		for i, v := range a[1:] {
			if v < big.IntSlice[0] {
				small.push(-v)
			} else {
				big.push(v)
			}
			if len(big.IntSlice)-1 > len(small.IntSlice) {
				small.push(-big.pop())
			} else if len(small.IntSlice) > len(big.IntSlice) {
				big.push(-small.pop())
			}
			if i&1 == 1 {
				medians = append(medians, big.IntSlice[0])
			}
		}
		return medians
	}

	// 离线做法，使用链表
	dynamicMediansOffline := func(a []int) []int {
		panic("TODO")
	}

	_ = []interface{}{dynamicMedians, dynamicMediansOffline}
}
