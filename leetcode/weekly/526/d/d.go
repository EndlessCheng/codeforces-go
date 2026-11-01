package main

import (
	"math"
	"math/bits"
)

// https://space.bilibili.com/206214
func minActivationCost1(edges [][]int, values []int) int {
	type edge struct{ to, c int }
	g := make([][]edge, len(values))
	for _, e := range edges {
		x, y, c := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, c})
		g[y] = append(g[y], edge{x, c})
	}

	const inf = math.MaxInt / 4
	ans := inf

	// pOnes 为根到 x 的路径异或和的二进制中的 1 的个数的奇偶性
	// 返回从 x 到子树 y 中节点的最小开销和（按照 pOnes 的奇偶性分别计算）
	var dfs func(int, int, int) [2]int
	dfs = func(x, fa, pOnes int) [2]int {
		pVal := bits.OnesCount(uint(values[x])) % 2
		pOnes ^= pVal
		minSum := [2]int{inf, inf}
		minSum[pOnes] = 0 // x 作为路径的端点

		for _, e := range g[x] {
			y := e.to
			if y == fa {
				continue
			}
			mn := dfs(y, x, pOnes)
			mn[0] += e.c
			mn[1] += e.c
			// 现在 mn 是从 x 到子树 y 中节点的最小开销和
			// 拼接从 x 到两棵子树中的路径（拼接后，路径在 x 拐弯）
			ans = min(ans, minSum[0]+mn[pVal], minSum[1]+mn[pVal^1])
			minSum[0] = min(minSum[0], mn[0])
			minSum[1] = min(minSum[1], mn[1])
		}

		return minSum
	}
	dfs(0, -1, 0)

	if ans < inf {
		return ans
	}
	return -1
}

func minActivationCost2(edges [][]int, values []int) int {
	const inf = math.MaxInt / 2
	ans := inf
	groups := make([][]int, len(values))

	for _, e := range edges {
		x, y, c := e[0], e[1], e[2]
		vx := bits.OnesCount(uint(values[x])) % 2
		vy := bits.OnesCount(uint(values[y])) % 2
		if vx == vy {
			ans = min(ans, c)
		} else if vx == 0 { // 此时 vy = 1
			// 三节点路径的点权必须是 1,0,1，中间的点权必须是 0
			groups[x] = append(groups[x], c)
		} else { // 此时 vy = 0 且 vx = 1
			groups[y] = append(groups[y], c)
		}
	}

	// 枚举三节点路径的中间节点
	for _, cs := range groups {
		// 枚举相邻边权，计算两数之和的最小值（枚举右维护左）
		mn := inf
		for _, c := range cs {
			ans = min(ans, mn+c)
			mn = min(mn, c)
		}
	}

	if ans < inf {
		return ans
	}
	return -1
}

func minActivationCost(edges [][]int, values []int) int {
	const inf = math.MaxInt / 2
	ans := inf
	mn := make([]int, len(values))
	for i := range mn {
		mn[i] = inf
	}

	for _, e := range edges {
		x, y, c := e[0], e[1], e[2]
		vx := bits.OnesCount(uint(values[x])) % 2
		vy := bits.OnesCount(uint(values[y])) % 2
		if vx == vy {
			ans = min(ans, c)
		} else if vx == 0 { // 此时 vy = 1
			ans = min(ans, mn[x]+c)
			mn[x] = min(mn[x], c)
		} else { // 此时 vy = 0 且 vx = 1
			ans = min(ans, mn[y]+c)
			mn[y] = min(mn[y], c)
		}
	}

	if ans < inf {
		return ans
	}
	return -1
}
