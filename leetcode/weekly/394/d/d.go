package main

import (
	"container/heap"
	"math"
)

// https://space.bilibili.com/206214
func findAnswer(n int, edges [][]int) []bool {
	type edge struct{ to, w, i int }
	g := make([][]edge, n)
	for i, e := range edges {
		x, y, w := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, w, i})
		g[y] = append(g[y], edge{x, w, i})
	}

	// Dijkstra 算法模板
	dis := make([]int, n)
	for i := 1; i < n; i++ {
		dis[i] = math.MaxInt
	}
	h := hp{{}}
	for len(h) > 0 {
		p := heap.Pop(&h).(pair)
		x := p.x
		if p.dis > dis[x] {
			continue
		}
		for _, e := range g[x] {
			y := e.to
			newD := p.dis + e.w
			if newD < dis[y] {
				dis[y] = newD
				heap.Push(&h, pair{newD, y})
			}
		}
	}

	ans := make([]bool, len(edges))
	// 图不连通
	if dis[n-1] == math.MaxInt {
		return ans
	}

	// 从终点出发 DFS
	vis := make([]bool, n)
	var dfs func(int)
	dfs = func(y int) {
		vis[y] = true
		for _, e := range g[y] {
			x := e.to
			if dis[x]+e.w != dis[y] {
				continue
			}
			ans[e.i] = true
			if !vis[x] {
				dfs(x)
			}
		}
	}
	dfs(n - 1)
	return ans
}

func findAnswer2(n int, edges [][]int) []bool {
	type edge struct{ to, w, i int }
	g := make([][]edge, n)
	for i, e := range edges {
		x, y, w := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, w, i})
		g[y] = append(g[y], edge{x, w, i})
	}

	dis := make([]int, n)
	for i := 1; i < n; i++ {
		dis[i] = math.MaxInt
	}
	h := hp{{}}
	for len(h) > 0 {
		p := heap.Pop(&h).(pair)
		x := p.x
		if p.dis > dis[x] {
			continue
		}
		for _, e := range g[x] {
			y := e.to
			newD := p.dis + e.w
			if newD < dis[y] {
				dis[y] = newD
				heap.Push(&h, pair{newD, y})
			}
		}
	}

	ans := make([]bool, len(edges))
	// 图不连通
	if dis[n-1] == math.MaxInt {
		return ans
	}

	// 从终点出发 BFS
	vis := make([]bool, n)
	vis[n-1] = true
	q := []int{n - 1}
	for len(q) > 0 {
		y := q[0]
		q = q[1:]
		for _, e := range g[y] {
			x := e.to
			if dis[x]+e.w != dis[y] {
				continue
			}
			ans[e.i] = true
			if !vis[x] {
				vis[x] = true
				q = append(q, x)
			}
		}
	}
	return ans
}

type pair struct{ dis, x int }
type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
