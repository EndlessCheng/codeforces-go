package main

import "slices"

// https://space.bilibili.com/206214
func placedCoins(edges [][]int, cost []int) []int64 {
	n := len(cost)
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	ans := make([]int64, n)
	var dfs func(int, int) []int
	dfs = func(x, fa int) []int {
		a := []int{cost[x]}
		for _, y := range g[x] {
			if y != fa {
				a = append(a, dfs(y, x)...)
			}
		}

		slices.Sort(a)
		m := len(a)
		if m < 3 {
			ans[x] = 1
		} else {
			ans[x] = int64(max(a[m-3]*a[m-2]*a[m-1], a[0]*a[1]*a[m-1], 0))
		}
		if m > 5 {
			a = append(a[:2], a[m-3:]...)
		}
		return a
	}
	dfs(0, -1)
	return ans
}
