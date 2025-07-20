package main

import (
	"math"
	"sort"
)

// https://space.bilibili.com/206214
func findMaxPathScore(edges [][]int, online []bool, k int64) int {
	n := len(online)
	type edge struct{ to, wt int }
	g := make([][]edge, n)
	maxWt := 0
	for _, e := range edges {
		x, y, wt := e[0], e[1], e[2]
		if online[x] && online[y] {
			g[x] = append(g[x], edge{y, wt})
			maxWt = max(maxWt, wt)
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
