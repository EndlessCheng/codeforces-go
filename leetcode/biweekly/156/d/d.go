package main

import "math"

// https://space.bilibili.com/206214
func subtreeInversionSum(edges [][]int, nums []int, k int) int64 {
	n := len(nums)
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	memo := make([][][2]int, n)
	for i := range memo {
		memo[i] = make([][2]int, k)
		for j := range memo[i] {
			for p := range memo[i][j] {
				memo[i][j][p] = math.MinInt
			}
		}
	}
	var dfs func(int, int, int, int) int
	dfs = func(x, fa, cd, parity int) int {
		p := &memo[x][cd][parity]
		if *p != math.MinInt {
			return *p
		}

		// 不反转
		res := nums[x] * (1 - parity*2)
		for _, y := range g[x] {
			if y != fa {
				res += dfs(y, x, max(cd-1, 0), parity)
			}
		}

		// 反转
		if cd == 0 {
			s := nums[x] * (parity*2 - 1)
			for _, y := range g[x] {
				if y != fa {
					s += dfs(y, x, k-1, parity^1) // 重置 CD
				}
			}
			res = max(res, s)
		}

		*p = res
		return res
	}
	return int64(dfs(0, -1, 0, 0))
}
