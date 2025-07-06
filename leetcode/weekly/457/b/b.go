package main

import (
	"container/heap"
	"math"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func processQueries1(c int, connections [][]int, queries [][]int) (ans []int) {
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

func processQueries(c int, connections [][]int, queries [][]int) []int {
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
	cc := 0

	var dfs func(int)
	dfs = func(x int) {
		belong[x] = cc
		for _, y := range g[x] {
			if belong[y] < 0 {
				dfs(y)
			}
		}
	}
	for i := 1; i <= c; i++ {
		if belong[i] < 0 {
			dfs(i)
			cc++
		}
	}

	offlineTime := make([]int, c+1)
	for i := range offlineTime {
		offlineTime[i] = math.MaxInt
	}
	q1 := 0
	for i, q := range slices.Backward(queries) {
		if q[0] == 2 {
			offlineTime[q[1]] = i // 记录最早离线时间
		} else {
			q1++
		}
	}

	// 维护每个连通块的在线电站的最小编号
	mn := make([]int, cc)
	for i := range mn {
		mn[i] = math.MaxInt
	}
	for i := 1; i <= c; i++ {
		if offlineTime[i] == math.MaxInt { // 最终仍然在线
			j := belong[i]
			mn[j] = min(mn[j], i)
		}
	}

	ans := make([]int, q1)
	for i, q := range slices.Backward(queries) {
		x := q[1]
		j := belong[x]
		if q[0] == 2 {
			if offlineTime[x] == i { // 变回在线
				mn[j] = min(mn[j], x)
			}
		} else {
			q1--
			if i < offlineTime[x] { // 已经在线（写 < 或者 <= 都可以）
				ans[q1] = x
			} else if mn[j] != math.MaxInt {
				ans[q1] = mn[j]
			} else {
				ans[q1] = -1
			}
		}
	}
	return ans
}
