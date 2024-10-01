package main

import (
	"container/heap"
	"math"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func minNumberOfSeconds(mountainHeight int, workerTimes []int) int64 {
	maxT := slices.Max(workerTimes)
	h := (mountainHeight-1)/len(workerTimes) + 1
	ans := 1 + sort.Search(maxT*h*(h+1)/2-1, func(m int) bool {
		m++
		leftH := mountainHeight
		for _, t := range workerTimes {
			leftH -= (int(math.Sqrt(float64(m/t*8+1))) - 1) / 2
			if leftH <= 0 {
				return true
			}
		}
		return false
	})
	return int64(ans)
}

func minNumberOfSeconds2(mountainHeight int, workerTimes []int) int64 {
	h := make(hp, len(workerTimes))
	for i, t := range workerTimes {
		h[i] = worker{t, t, t}
	}
	heap.Init(&h)

	ans := 0
	for ; mountainHeight > 0; mountainHeight-- {
		ans = h[0].nxt
		h[0].delta += h[0].base
		h[0].nxt += h[0].delta
		heap.Fix(&h, 0)
	}
	return int64(ans)
}

type worker struct{ nxt, delta, base int }
type hp []worker

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].nxt < h[j].nxt }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (hp) Push(any)             {}
func (hp) Pop() (_ any)         { return }
