package main

import "math"

// https://space.bilibili.com/206214
func maximumValueSum(nums []int, k int, _ [][]int) int64 {
	f0, f1 := 0, math.MinInt
	for _, x := range nums {
		f0, f1 = max(f0+x, f1+(x^k)), max(f1+x, f0+(x^k))
	}
	return int64(f0)
}

func maximumValueSum2(nums []int, k int, edges [][]int) int64 {
	n := len(nums)
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	var dfs func(int, int) (int, int)
	dfs = func(x, fa int) (int, int) {
		f0, f1 := 0, math.MinInt
		for _, y := range g[x] {
			if y != fa {
				r0, r1 := dfs(y, x)
				f0, f1 = max(f0+r0, f1+r1), max(f1+r0, f0+r1)
			}
		}
		return max(f0+nums[x], f1+(nums[x]^k)), max(f1+nums[x], f0+(nums[x]^k))
	}
	ans, _ := dfs(0, -1)
	return int64(ans)
}
