package main

import "math/bits"

// https://space.bilibili.com/206214
func maxProfit(n int, edges [][]int, score []int) int {
	// 记录每个节点的先修课（直接前驱）
	pre := make([]int, n)
	for _, e := range edges {
		pre[e[1]] |= 1 << e[0]
	}

	memo := make([]int, 1<<n)
	var dfs func(s int) int
	dfs = func(s int) (res int) {
		m := &memo[s]
		if *m > 0 {
			return *m
		}
		defer func() { *m = res }()
		i := bits.OnesCount(uint(s)) // 已学课程数
		// 枚举还没学过的课程 j，且 j 的所有先修课都学完了
		for j, p := range pre {
			if s>>j&1 == 0 && s|p == s {
				res = max(res, dfs(s|1<<j)+score[j]*(i+1))
			}
		}
		return
	}
	return dfs(0)
}

func maxProfit2(n int, edges [][]int, score []int) int {
	// 记录每个节点的先修课（直接前驱）
	pre := make([]int, n)
	for _, e := range edges {
		pre[e[1]] |= 1 << e[0]
	}

	u := 1 << n
	f := make([]int, u)

	for s := u - 2; s >= 0; s-- {
		i := bits.OnesCount(uint(s)) // 已学课程数
		// 枚举还没学过的课程 j，且 j 的所有先修课都学完了
		for j, p := range pre {
			if s>>j&1 == 0 && s|p == s {
				f[s] = max(f[s], f[s|1<<j]+score[j]*(i+1))
			}
		}
	}
	return f[0]
}
