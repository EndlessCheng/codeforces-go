package main

import (
	"slices"
)

func smallestStringWithSwaps(s string, pairs [][]int) string {
	n := len(s)
	// 初始化并查集
	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	// 合并
	for _, p := range pairs {
		fa[find(p[0])] = find(p[1])
	}

	// 把同一个连通块中的下标分到同一组
	idxGroups := make([][]int, n)
	for i := range n {
		rt := find(i)
		idxGroups[rt] = append(idxGroups[rt], i)
	}

	ans := make([]byte, n)
	for _, idx := range idxGroups {
		slices.Sort(idx) // 下标从小到大排序

		t := make([]byte, len(idx))
		for i, j := range idx {
			t[i] = s[j]
		}
		slices.Sort(t) // 字母从小到大排序

		// 小的字母填入小的下标，大的字母填入大的下标
		for i, ch := range t {
			ans[idx[i]] = ch
		}
	}
	return string(ans)
}

func smallestStringWithSwaps1(s string, pairs [][]int) string {
	n := len(s)
	g := make([][]int, n)
	for _, p := range pairs {
		x, y := p[0], p[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	vis := make([]bool, n)
	idx := []int{}

	var dfs func(int)
	dfs = func(x int) {
		vis[x] = true
		idx = append(idx, x) // 收集这个连通块的下标
		for _, y := range g[x] {
			if !vis[y] {
				dfs(y)
			}
		}
	}

	ans := make([]byte, n)
	for start, b := range vis {
		if b {
			continue
		}

		idx = idx[:0]
		dfs(start)
		slices.Sort(idx) // 下标从小到大排序

		t := make([]byte, len(idx))
		for i, j := range idx {
			t[i] = s[j]
		}
		slices.Sort(t) // 字母从小到大排序

		// 小的字母填入小的下标，大的字母填入大的下标
		for i, ch := range t {
			ans[idx[i]] = ch
		}
	}
	return string(ans)
}
