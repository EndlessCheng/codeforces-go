package main

import (
	"container/heap"
	"math"
)

// github.com/EndlessCheng/codeforces-go
const inf = math.MaxInt / 3 // 防止计算 d1[x] + d2[x] + d3[x] 时溢出

type edge struct{ to, wt int }

func minimumWeight(n int, edges [][]int, src1, src2, dest int) int64 {
	g := make([][]edge, n)
	rg := make([][]edge, n) // 反图，用于计算从 dest 出发的最短路
	for _, e := range edges {
		x, y, wt := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, wt})
		rg[y] = append(rg[y], edge{x, wt})
	}

	d1 := dijkstra(g, src1)
	d2 := dijkstra(g, src2)
	d3 := dijkstra(rg, dest)

	ans := inf
	for x := range n { // 枚举相遇点 x
		ans = min(ans, d1[x]+d2[x]+d3[x])
	}

	if ans == inf {
		return -1
	}
	return int64(ans)
}

func dijkstra(g [][]edge, start int) []int {
	dis := make([]int, len(g))
	for i := range dis {
		dis[i] = inf
	}
	dis[start] = 0 // 起点到自己的距离是 0
	// 堆中保存 (起点到节点 x 的最短路长度，节点 x)
	h := &hp{{0, start}}

	for h.Len() > 0 {
		p := heap.Pop(h).(pair)
		disX, x := p.dis, p.x
		if disX > dis[x] { // x 之前出堆过
			continue
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

	return dis
}

type pair struct{ dis, x int }
type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
