package main

import (
	"math/bits"
	"slices"
)

// https://space.bilibili.com/206214
func maxProfit(n int, edges [][]int, score []int) (ans int) {
	if len(edges) == 0 {
		slices.Sort(score)
		for i, s := range score {
			ans += s * (i + 1)
		}
		return
	}

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

// 超时
func maxProfit2(n int, edges [][]int, score []int) (ans int) {
	if len(edges) == 0 {
		slices.Sort(score)
		for i, s := range score {
			ans += s * (i + 1)
		}
		return
	}

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

func maxProfit3(n int, edges [][]int, score []int) (ans int) {
	if len(edges) == 0 {
		slices.Sort(score)
		for i, s := range score {
			ans += s * (i + 1)
		}
		return
	}

	// 记录每个节点的先修课（直接前驱）
	pre := make([]int, n)
	for _, e := range edges {
		pre[e[1]] |= 1 << e[0]
	}

	u := 1 << n
	f := make([]int, u)
	for s := 1; s < u; s++ {
		f[s] = -1
	}

	for s, fs := range f {
		if fs < 0 { // 不合法状态，比如已经学完后面的课程，但前面的课程还没学
			continue
		}
		i := bits.OnesCount(uint(s)) // 已学课程数
		// 枚举还没学过的课程 j，且 j 的所有先修课都学完了
		for cus, lb := u-1^s, 0; cus > 0; cus ^= lb {
			lb = cus & -cus
			j := bits.TrailingZeros(uint(lb))
			if s|pre[j] == s {
				newS := s | lb
				f[newS] = max(f[newS], fs+score[j]*(i+1))
			}
		}
	}
	return f[u-1]
}
