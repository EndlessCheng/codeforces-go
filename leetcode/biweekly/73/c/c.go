package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func getAncestors(n int, es [][]int) [][]int {
	g := make([][]int, n)
	for _, e := range es {
		g[e[1]] = append(g[e[1]], e[0])
	}
	vis := make([]int, n)
	ans := make([][]int, n)
	for i := range ans {
		a := []int{}
		var dfs func(int)
		dfs = func(v int) {
			vis[v] = i + 1
			a = append(a, v)
			for _, w := range g[v] {
				if vis[w] != i+1 {
					dfs(w)
				}
			}
		}
		dfs(i)
		a = a[1:] // 去掉 i 本身
		sort.Ints(a)
		ans[i] = a
	}
	return ans
}
