package copypasta

import (
	"container/heap"
	"sort"
)

// 下面这些都是最小堆

type hp struct{ sort.IntSlice }

//func (h hp) Less(i, j int) bool  { return h.IntSlice[i] > h.IntSlice[j] } // 最大堆
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
func (h *hp) push(v int)         { heap.Push(h, v) }
func (h *hp) pop() int           { return heap.Pop(h).(int) }
func (h hp) empty() bool         { return len(h.IntSlice) == 0 }
func (h hp) top() int            { return h.IntSlice[0] }
func (h hp) size() int           { return len(h.IntSlice) }
func (h *hp) pushPop(v int) int {
	if len(h.IntSlice) > 0 && v > h.IntSlice[0] { // 大根堆改成 v < h.IntSlice[0]
		v, h.IntSlice[0] = h.IntSlice[0], v
		heap.Fix(h, 0)
	}
	return v
}
func (h *hp) popPush(v int) int { t := h.IntSlice[0]; h.IntSlice[0] = v; heap.Fix(h, 0); return t } // h 需要非空

//

type hp64 []int64 // 自定义类型

func (h hp64) Len() int            { return len(h) }
func (h hp64) Less(i, j int) bool  { return h[i] < h[j] } // > 为最大堆
func (h hp64) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp64) Push(v interface{}) { *h = append(*h, v.(int64)) }
func (h *hp64) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func (h *hp64) push(v int64)       { heap.Push(h, v) }
func (h *hp64) pop() int64         { return heap.Pop(h).(int64) }
func (h hp64) empty() bool         { return len(h) == 0 }
func (h hp64) top() int64          { return h[0] }
func (h hp64) size() int           { return len(h) }
func (h *hp64) pushPop(v int64) int64 {
	if len(*h) > 0 && v > (*h)[0] { // 大根堆改成 v < (*h)[0]
		v, (*h)[0] = (*h)[0], v
		heap.Fix(h, 0)
	}
	return v
}
func (h *hp64) popPush(v int64) int64 { t := (*h)[0]; (*h)[0] = v; heap.Fix(h, 0); return t } // h 需要非空

//

func heapCollections() {
	// 对顶堆求动态中位数（一个数组前 2k+1 项的中位数）
	// https://www.luogu.com.cn/problem/P1168
	// LC295 https://leetcode-cn.com/problems/find-median-from-data-stream/
	dynamicMedians := func(a []int) []int {
		n := len(a)
		medians := make([]int, 0, (n+1)/2)
		small, big := hp{}, hp{}
		for i, v := range a {
			if small.size() == big.size() {
				big.push(-small.pushPop(-v))
			} else {
				small.push(-big.pushPop(v))
			}
			if i&1 == 0 {
				medians = append(medians, big.top())
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
