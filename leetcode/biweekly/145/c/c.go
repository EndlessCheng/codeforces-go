package main

import (
	"container/heap"
	"math"
	"strconv"
)

// https://space.bilibili.com/206214
const mx = 10000
var np = [mx]bool{1: true}

func init() {
	// 埃氏筛，标记每个数是否为合数（或者 1）
	for i := 2; i < mx; i++ {
		if !np[i] {
			for j := i * i; j < mx; j += i {
				np[j] = true // 合数
			}
		}
	}
}

func minOperations(n, m int) int {
	if !np[n] || !np[m] {
		return -1
	}
	lenN := len(strconv.Itoa(n))
	dis := make([]int, int(math.Pow10(lenN)))
	for i := range dis {
		dis[i] = math.MaxInt
	}
	dis[n] = n
	h := hp{{n, n}}
	for len(h) > 0 {
		top := heap.Pop(&h).(pair)
		x := top.x
		if x == m {
			return top.dis
		}
		disX := top.dis
		if disX > dis[x] {
			continue
		}
		pow10 := 1
		for v := x; v > 0; v /= 10 {
			d := v % 10
			if d > 0 { // 减少
				y := x - pow10
				newD := disX + y
				if np[y] && newD < dis[y] {
					dis[y] = newD
					heap.Push(&h, pair{newD, y})
				}
			}
			if d < 9 { // 增加
				y := x + pow10
				newD := disX + y
				if np[y] && newD < dis[y] {
					dis[y] = newD
					heap.Push(&h, pair{newD, y})
				}
			}
			pow10 *= 10
		}
	}
	return -1
}

type pair struct{ dis, x int }
type hp []pair
func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
