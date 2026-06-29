package main

import (
	"container/heap"
	"math"
)

// https://space.bilibili.com/206214
func minTimeMaxPower(n int, edges [][]int, power int, cost []int, source int, target int) []int64 {
	type edge struct{ to, t int }
	g := make([][]edge, n)
	for _, e := range edges {
		x, y, t := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, t})
	}

	dis := make([][]int, n)
	for i := range dis {
		dis[i] = make([]int, power+1)
		for j := range dis[i] {
			dis[i][j] = math.MaxInt
		}
	}
	dis[source][power] = 0
	h := hp{{0, source, power}}

	for len(h) > 0 {
		top := heap.Pop(&h).(tuple)
		d, x, rem := top.dis, top.x, top.rem
		if x == target {
			return []int64{int64(d), int64(rem)}
		}
		if d > dis[x][rem] || rem < cost[x] {
			continue
		}
		rem -= cost[x]
		for _, e := range g[x] {
			y := e.to
			newD := d + e.t
			if newD < dis[y][rem] {
				dis[y][rem] = newD
				heap.Push(&h, tuple{newD, y, rem})
			}
		}
	}

	return []int64{-1, -1}
}

// 最短路长度, 节点编号, 剩余电量
type tuple struct{ dis, x, rem int }
type hp []tuple

func (h hp) Len() int { return len(h) }
func (h hp) Less(i, j int) bool {
	a, b := h[i], h[j]
	return a.dis < b.dis || a.dis == b.dis && a.rem > b.rem
}
func (h hp) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)   { *h = append(*h, v.(tuple)) }
func (h *hp) Pop() (v any) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }

//

func minTimeMaxPower2(n int, edges [][]int, power int, cost []int, source int, target int) []int64 {
	type edge struct{ to, t int }
	g := make([][]edge, n)
	for _, e := range edges {
		x, y, t := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, t})
	}

	f := make([][]int, power+1)
	for i := range f {
		f[i] = make([]int, n)
		for j := range f[i] {
			f[i][j] = math.MaxInt
		}
	}
	f[power][source] = 0

	minDis, maxRem := math.MaxInt, -1
	for rem := power; rem >= 0; rem-- {
		if f[rem][target] < minDis {
			minDis, maxRem = f[rem][target], rem
		}
		for x, v := range f[rem] {
			if v == math.MaxInt || rem < cost[x] {
				continue
			}
			nxtRem := rem - cost[x]
			for _, e := range g[x] {
				f[nxtRem][e.to] = min(f[nxtRem][e.to], v+e.t) // 刷表法
			}
		}
	}

	if maxRem < 0 {
		return []int64{-1, -1}
	}
	return []int64{int64(minDis), int64(maxRem)}
}
