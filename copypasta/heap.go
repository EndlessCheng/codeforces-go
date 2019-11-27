package copypasta

import "sort"

// 最小堆
type intHeap struct {
	sort.IntSlice
}

//func (h *intHeap) Len() int           { return len(h.IntSlice) }
//func (h *intHeap) Less(i, j int) bool { return h.IntSlice[i] < h.IntSlice[j] } // > 为最大堆
//func (h *intHeap) Swap(i, j int)      { h.IntSlice[i], h.IntSlice[j] = h.IntSlice[j], h.IntSlice[i] }
func (h *intHeap) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *intHeap) Pop() (v interface{}) {
	n := len(h.IntSlice)
	h.IntSlice, v = h.IntSlice[:n-1], h.IntSlice[n-1]
	return
}

//

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
