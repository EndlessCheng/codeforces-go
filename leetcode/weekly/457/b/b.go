package main

import (
	"container/heap"
	"sort"
)

// https://space.bilibili.com/206214
func processQueries(c int, connections [][]int, queries [][]int) (ans []int) {
	g := make([][]int, c+1)
	for _, e := range connections {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	belong := make([]int, c+1)
	for i := range belong {
		belong[i] = -1
	}
	heaps := []hp{}
	var h hp

	var dfs func(int)
	dfs = func(x int) {
		belong[x] = len(heaps) // 记录节点 x 在哪个堆
		h.IntSlice = append(h.IntSlice, x)
		for _, y := range g[x] {
			if belong[y] < 0 {
				dfs(y)
			}
		}
	}
	for i := 1; i <= c; i++ {
		if belong[i] >= 0 {
			continue
		}
		h = hp{}
		dfs(i)
		heap.Init(&h)
		heaps = append(heaps, h)
	}

	offline := make([]bool, c+1)
	for _, q := range queries {
		x := q[1]
		if q[0] == 2 {
			offline[x] = true
			continue
		}
		if !offline[x] {
			ans = append(ans, x)
			continue
		}
		// 懒删除：取堆顶的时候，如果离线，才删除
		h := &heaps[belong[x]]
		for h.Len() > 0 && offline[h.IntSlice[0]] {
			heap.Pop(h)
		}
		if h.Len() > 0 {
			ans = append(ans, h.IntSlice[0])
		} else {
			ans = append(ans, -1)
		}
	}
	return
}

type hp struct{ sort.IntSlice }
func (h *hp) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
