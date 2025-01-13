package main

import "math"

// https://space.bilibili.com/206214
func maximumAmount(coins [][]int) int {
	n := len(coins[0])
	f := make([][3]int, n+1)
	for j := range f {
		f[j] = [3]int{math.MinInt / 2, math.MinInt / 2, math.MinInt / 2}
	}
	f[1] = [3]int{}
	for _, row := range coins {
		for j, x := range row {
			f[j+1][2] = max(f[j][2]+x, f[j+1][2]+x, f[j][1], f[j+1][1])
			f[j+1][1] = max(f[j][1]+x, f[j+1][1]+x, f[j][0], f[j+1][0])
			f[j+1][0] = max(f[j][0], f[j+1][0]) + x
		}
	}
	return f[n][2]
}

func maximumAmount2(coins [][]int) int {
	m, n := len(coins), len(coins[0])
	f := make([][][3]int, m+1)
	for i := range f {
		f[i] = make([][3]int, n+1)
	}
	for j := range f[0] {
		f[0][j] = [3]int{math.MinInt / 2, math.MinInt / 2, math.MinInt / 2}
	}
	f[0][1] = [3]int{}
	for i, row := range coins {
		f[i+1][0] = [3]int{math.MinInt / 2, math.MinInt / 2, math.MinInt / 2}
		for j, x := range row {
			f[i+1][j+1][0] = max(f[i+1][j][0], f[i][j+1][0]) + x
			f[i+1][j+1][1] = max(f[i+1][j][1]+x, f[i][j+1][1]+x, f[i+1][j][0], f[i][j+1][0])
			f[i+1][j+1][2] = max(f[i+1][j][2]+x, f[i][j+1][2]+x, f[i+1][j][1], f[i][j+1][1])
		}
	}
	return f[m][n][2]
}

func maximumAmount1(coins [][]int) int {
	m, n := len(coins), len(coins[0])
	memo := make([][][3]int, m)
	for i := range memo {
		memo[i] = make([][3]int, n)
		for j := range memo[i] {
			for k := range memo[i][j] {
				memo[i][j][k] = math.MinInt
			}
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, j, k int) int {
		if i < 0 || j < 0 {
			return math.MinInt
		}
		x := coins[i][j]
		if i == 0 && j == 0 {
			if k == 0 {
				return x
			}
			return max(x, 0)
		}
		p := &memo[i][j][k]
		if *p != math.MinInt {
			return *p
		}
		res := max(dfs(i-1, j, k), dfs(i, j-1, k)) + x
		if x < 0 && k > 0 {
			res = max(res, dfs(i-1, j, k-1), dfs(i, j-1, k-1))
		}
		*p = res
		return res
	}
	return dfs(m-1, n-1, 2)
}
