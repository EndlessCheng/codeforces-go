package main

import (
	"container/heap"
	"math"
	"math/bits"
)

// https://space.bilibili.com/206214
func minTime(n, k, m int, time []int, mul []float64) float64 {
	u := 1 << n
	// 计算每个 time 子集的最大值
	maxTime := make([]int, u)
	for i, t := range time {
		highBit := 1 << i
		for mask, mx := range maxTime[:highBit] {
			maxTime[highBit|mask] = max(mx, t)
		}
	}
	// 把 maxTime 中的大小大于 k 的集合改为 inf
	for i := range maxTime {
		if bits.OnesCount(uint(i)) > k {
			maxTime[i] = math.MaxInt
		}
	}

	dis := make([][][2]float64, m)
	for i := range dis {
		dis[i] = make([][2]float64, u)
		for j := range dis[i] {
			dis[i][j] = [2]float64{math.MaxFloat64, math.MaxFloat64}
		}
	}
	h := hp{}
	push := func(d float64, stage, mask int, state uint8) {
		if d < dis[stage][mask][state] {
			dis[stage][mask][state] = d
			heap.Push(&h, tuple{d, stage, mask, state})
		}
	}

	push(0, 0, u-1, 0) // 起点

	for len(h) > 0 {
		top := heap.Pop(&h).(tuple)
		d := top.dis
		stage := top.stage
		left := top.mask // 剩余没有过河的人
		state := top.state
		if left == 0 { // 所有人都过河了
			return d
		}
		if d > dis[stage][left][state] {
			continue
		}
		if state == 0 {
			// 枚举 sub 这群人坐一艘船
			for sub := left; sub > 0; sub = (sub - 1) & left {
				if maxTime[sub] != math.MaxInt {
					cost := float64(maxTime[sub]) * mul[stage]
					push(d+cost, (stage+int(cost))%m, left^sub, 1)
				}
			}
		} else {
			// 枚举回来的人
			for s, lb := u-1^left, 0; s > 0; s ^= lb {
				lb = s & -s
				cost := float64(maxTime[lb]) * mul[stage]
				push(d+cost, (stage+int(cost))%m, left^lb, 0)
			}
		}
	}
	return -1
}

type tuple struct {
	dis         float64
	stage, mask int
	state       uint8 // 状态机：0 未过河，1 已过河
}
type hp []tuple

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(tuple)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
