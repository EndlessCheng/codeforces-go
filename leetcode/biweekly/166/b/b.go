package main

import "math"

// https://space.bilibili.com/206214
func climbStairs1(n int, costs []int) int {
	memo := make([]int, n+1)
	var dfs func(int) int
	dfs = func(i int) int {
		if i == 0 {
			return 0
		}
		p := &memo[i]
		if *p != 0 {
			return *p
		}
		res := math.MaxInt
		for j := max(i-3, 0); j < i; j++ {
			res = min(res, dfs(j)+(i-j)*(i-j))
		}
		res += costs[i-1]
		*p = res
		return res
	}
	return dfs(n)
}

func climbStairs2(n int, costs []int) int {
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		res := math.MaxInt
		for j := max(i-3, 0); j < i; j++ {
			res = min(res, f[j]+(i-j)*(i-j))
		}
		f[i] = res + costs[i-1]
	}
	return f[n]
}

func climbStairs(n int, costs []int) int {
	var f0, f1, f2 int
	for _, c := range costs {
		f0, f1, f2 = f1, f2, min(f0+9, f1+4, f2+1)+c
	}
	return f2
}
