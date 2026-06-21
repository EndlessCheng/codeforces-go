package main

import "math"

// https://space.bilibili.com/206214
func finishTime(n int, edges [][]int, baseTime []int) int64 {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y) // 题目保证 x 是 y 的父节点
	}

	var dfs func(int) int
	dfs = func(x int) int {
		if g[x] == nil { // x 是叶子
			return baseTime[x]
		}
		earliest, latest := math.MaxInt, 0
		for _, y := range g[x] {
			t := dfs(y)
			earliest = min(earliest, t)
			latest = max(latest, t)
		}
		return latest*2 - earliest + baseTime[x]
	}

	return int64(dfs(0))
}

// todo 添加爆 LL 的数据
// todo hack 边排序
