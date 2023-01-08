package main

import (
	"container/heap"
	"sort"
)

// https://space.bilibili.com/206214
func findCrossingTime(n, k int, time [][]int) (cur int) {
	sort.SliceStable(time, func(i, j int) bool {
		a, b := time[i], time[j]
		return a[0]+a[2] < b[0]+b[2]
	})
	waitL, waitR := make(hp, k), hp{}
	for i := range waitL {
		waitL[i].i = k - 1 - i
	}
	workL, workR := hp2{}, hp2{}
	for n > 0 {
		for workL.Len() > 0 && workL[0].t <= cur {
			heap.Push(&waitL, heap.Pop(&workL)) // 左边完成放箱
		}
		for workR.Len() > 0 && workR[0].t <= cur {
			heap.Push(&waitR, heap.Pop(&workR)) // 右边完成搬箱
		}
		if waitR.Len() > 0 && waitR[0].t <= cur { // 右边过桥
			p := heap.Pop(&waitR).(pair)
			cur += time[p.i][2]
			p.t = cur + time[p.i][3] // 放箱，记录完成时间
			heap.Push(&workL, p)
		} else if waitL.Len() > 0 && waitL[0].t <= cur { // 左边过桥
			p := heap.Pop(&waitL).(pair)
			cur += time[p.i][0]
			p.t = cur + time[p.i][1] // 搬箱，记录完成时间
			heap.Push(&workR, p)
			n--
		} else if workL.Len() == 0 { // cur 过小，找个最小的放箱/搬箱完成时间来更新 cur
			cur = workR[0].t
		} else if workR.Len() == 0 {
			cur = workL[0].t
		} else {
			cur = min(workL[0].t, workR[0].t)
		}
	}
	for workR.Len() > 0 {
		p := heap.Pop(&workR).(pair)       // 右边完成搬箱
		cur = max(p.t, cur) + time[p.i][2] // 过桥
	}
	return cur // 最后一个过桥的时间
}

type pair struct{ i, t int }
type hp []pair
func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].i > h[j].i }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

type hp2 []pair
func (h hp2) Len() int            { return len(h) }
func (h hp2) Less(i, j int) bool  { return h[i].t < h[j].t }
func (h hp2) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp2) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp2) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func min(a, b int) int { if b < a { return b }; return a }
func max(a, b int) int { if b > a { return b }; return a }
