package main

import (
	"math"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func findMaxPathScore1(edges [][]int, online []bool, k int64) int {
	n := len(online)
	type edge struct{ to, wt int }
	g := make([][]edge, n)
	maxWt := -1
	for _, e := range edges {
		x, y, wt := e[0], e[1], e[2]
		if online[x] && online[y] {
			g[x] = append(g[x], edge{y, wt})
			if x == 0 {
				maxWt = max(maxWt, wt)
			}
		}
	}

	memo := make([]int, n)
	// 二分无法到达 n-1 的最小 lower，那么减一后，就是可以到达 n-1 的最大 lower
	ans := sort.Search(maxWt+1, func(lower int) bool {
		for i := range memo {
			memo[i] = -1
		}
		var dfs func(int) int
		dfs = func(x int) int {
			if x == n-1 { // 到达终点
				return 0
			}
			p := &memo[x]
			if *p != -1 { // 之前计算过
				return *p
			}
			res := math.MaxInt / 2 // 防止加法溢出
			for _, e := range g[x] {
				y := e.to
				if e.wt >= lower {
					res = min(res, dfs(y)+e.wt)
				}
			}
			*p = res // 记忆化
			return res
		}
		return dfs(0) > int(k)
	}) - 1
	return ans
}

func findMaxPathScore(edges [][]int, online []bool, k int64) int {
	n := len(online)
	type edge struct{ to, wt int }
	g := make([][]edge, n)
	deg := make([]int, n)
	maxWt := 0
	for _, e := range edges {
		x, y, wt := e[0], e[1], e[2]
		if online[x] && online[y] {
			g[x] = append(g[x], edge{y, wt})
			deg[y]++
			maxWt = max(maxWt, wt)
		}
	}

	// 先清理无法从 0 到达的边，比如 2 -> 0 
	q := []int{}
	for i := 1; i < n; i++ {
		if deg[i] == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		for _, e := range g[x] {
			y := e.to
			deg[y]--
			if deg[y] == 0 && y > 0 {
				q = append(q, y)
			}
		}
	}

	f := make([]int, n)
	ans := sort.Search(maxWt+1, func(lower int) bool {
		deg := slices.Clone(deg)
		for i := 1; i < n; i++ {
			f[i] = math.MaxInt / 2
		}

		q := []int{0}
		for len(q) > 0 {
			x := q[0]
			if x == n-1 {
				return f[x] > int(k)
			}
			q = q[1:]
			for _, e := range g[x] {
				y := e.to
				wt := e.wt
				if wt >= lower {
					f[y] = min(f[y], f[x]+wt)
				}
				deg[y]--
				if deg[y] == 0 {
					q = append(q, y)
				}
			}
		}
		return true
	}) - 1
	return ans
}
