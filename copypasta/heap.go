package copypasta

import (
	. "container/heap"
	"sort"
)

// 下面这些都是最小堆
// h.top() 即 h.IntSlice[0] 或 (*h)[0] （注意判断非空）

type hp struct{ sort.IntSlice }

//func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] } // 最大堆
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() (v interface{}) {
	n := len(h.IntSlice)
	h.IntSlice, v = h.IntSlice[:n-1], h.IntSlice[n-1]
	return
}

//

type hp64 []int64 // 自定义类型

func (h hp64) Len() int              { return len(h) }
func (h hp64) Less(i, j int) bool    { return h[i] < h[j] } // > 为最大堆
func (h hp64) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp64) Push(v interface{})   { *h = append(*h, v.(int64)) }
func (h *hp64) Pop() (v interface{}) { n := len(*h); *h, v = (*h)[:n-1], (*h)[n-1]; return }

//

// see graph.shortestPathDijkstra
type hPair struct {
	x int64
	y int
}
type pairHeap []hPair

func (h pairHeap) Len() int              { return len(h) }
func (h pairHeap) Less(i, j int) bool    { return h[i].x < h[j].x || h[i].x == h[j].x && h[i].y < h[j].y }
func (h pairHeap) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *pairHeap) Push(v interface{})   { *h = append(*h, v.(hPair)) }
func (h *pairHeap) Pop() (v interface{}) { n := len(*h); *h, v = (*h)[:n-1], (*h)[n-1]; return }

//

func heapCollections() {
	// 对顶堆求动态中位数（一个数组前 2k+1 项的中位数）
	// https://www.luogu.com.cn/problem/P1168
	dynamicMedians := func(a []int) []int {
		n := len(a)
		medians := make([]int, 1, (n+1)/2)
		medians[0] = a[0]
		small, big := &hp{}, &hp{}
		Push(big, a[0]) // 下面保证 big.size() == small.size() || big.size()-1 == small.size()
		for i, v := range a[1:] {
			if v < big.IntSlice[0] {
				Push(small, -v)
			} else {
				Push(big, v)
			}
			if len(big.IntSlice)-1 > len(small.IntSlice) {
				Push(small, -Pop(big).(int))
			} else if len(small.IntSlice) > len(big.IntSlice) {
				Push(big, -Pop(small).(int))
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
