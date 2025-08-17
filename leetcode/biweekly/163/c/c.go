package main

import (
	"container/heap"
	"math"
)

// https://space.bilibili.com/206214
func minCost(n int, edges [][]int) int {
	type edge struct{ to, wt int }
	g := make([][]edge, n) // 邻接表
	for _, e := range edges {
		x, y, wt := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, wt})
		g[y] = append(g[y], edge{x, wt * 2}) // 反转边
	}

	dis := make([]int, n)
	for i := range dis {
		dis[i] = math.MaxInt
	}
	dis[0] = 0 // 起点到自己的距离是 0
	// 堆中保存 (起点到节点 x 的最短路长度，节点 x)
	h := &hp{{}}

	for h.Len() > 0 {
		p := heap.Pop(h).(pair)
		disX, x := p.dis, p.x
		if disX > dis[x] { // x 之前出堆过
			continue
		}
		if x == n-1 { // 到达终点
			return disX
		}
		for _, e := range g[x] {
			y := e.to
			newDisY := disX + e.wt
			if newDisY < dis[y] {
				dis[y] = newDisY // 更新 x 的邻居的最短路
				// 懒更新堆：只插入数据，不更新堆中数据
				// 相同节点可能有多个不同的 newDisY，除了最小的 newDisY，其余值都会触发上面的 continue
				heap.Push(h, pair{newDisY, y})
			}
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
