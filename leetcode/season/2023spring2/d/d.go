package main

import (
	"sort"
	"strings"
)

// https://space.bilibili.com/206214
func evolutionaryRecord(parents []int) string {
	n := len(parents)
	g := make([][]int, n)
	for w := 1; w < n; w++ {
		p := parents[w]
		g[p] = append(g[p], w) // 建树
	}

	var dfs func(int) string
	dfs = func(v int) string {
		a := make([]string, len(g[v]))
		for i, w := range g[v] {
			a[i] = dfs(w)
		}
		sort.Strings(a)
		return "0" + strings.Join(a, "") + "1"
	}
	return strings.TrimRight(dfs(0)[1:], "1") // 去掉根节点以及返回根节点的路径
}
