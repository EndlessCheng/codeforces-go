package main

import (
	"container/heap"
	"math"
)

// https://space.bilibili.com/206214
func shortestPath(n int, edges [][]int, labels string, k int) int {
	type edge struct{ to, w int }
	g := make([][]edge, n)
	for _, e := range edges {
		x, y, w := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, w})
	}

	dis := make([][]int, n)
	for i := range dis {
		dis[i] = make([]int, k+1)
		for j := range dis[i] {
			dis[i][j] = math.MaxInt
		}
	}
	h := hp{}
	add := func(x, y, d int) {
		if d < dis[x][y] {
			dis[x][y] = d
			heap.Push(&h, tuple{d, x, y})
		}
	}

	add(0, 1, 0)
	for len(h) > 0 {
		top := heap.Pop(&h).(tuple)
		d := top.dis
		x, cnt := top.x, top.cnt
		if x == n-1 {
			return d
		}
		if d > dis[x][cnt] {
			continue
		}
		for _, e := range g[x] {
			y := e.to
			if labels[y] != labels[x] {
				add(y, 1, d+e.w)
			} else if cnt+1 <= k {
				add(y, cnt+1, d+e.w)
			}
		}
	}
	return -1
}

// 最短路长度, 节点编号, 最后连续相同字母个数
type tuple struct{ dis, x, cnt int }
type hp []tuple

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(tuple)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
