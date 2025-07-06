package main

import "math"

// https://space.bilibili.com/206214
func minCost1(m, n int, waitCost [][]int) int64 {
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
	}
	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		if i < 0 || j < 0 {
			return math.MaxInt
		}
		if i == 0 && j == 0 {
			return 1 // 起点只有 1 的进入成本，不需要等待
		}
		p := &memo[i][j]
		if *p == 0 {
			*p = min(dfs(i, j-1), dfs(i-1, j)) + waitCost[i][j] + (i+1)*(j+1)
		}
		return *p
	}
	return int64(dfs(m-1, n-1) - waitCost[m-1][n-1]) // 终点不需要等待
}

func minCost2(m, n int, waitCost [][]int) int64 {
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	f[0][1] = -waitCost[0][0] // 计算 f[1][1] 的时候抵消掉
	for j := 2; j <= n; j++ {
		f[0][j] = math.MaxInt
	}
	for i, row := range waitCost {
		f[i+1][0] = math.MaxInt
		for j, c := range row {
			f[i+1][j+1] = min(f[i+1][j], f[i][j+1]) + c + (i+1)*(j+1)
		}
	}
	return int64(f[m][n] - waitCost[m-1][n-1])
}

func minCost3(m, n int, waitCost [][]int) int64 {
	f := make([]int, n+1)
	for j := range f {
		f[j] = math.MaxInt
	}
	f[1] = -waitCost[0][0]
	for i, row := range waitCost {
		for j, c := range row {
			f[j+1] = min(f[j], f[j+1]) + c + (i+1)*(j+1)
		}
	}
	return int64(f[n] - waitCost[m-1][n-1])
}

func minCost(m, n int, f [][]int) int64 {
	f[0][0] = 1
	f[m-1][n-1] = 0
	for j := 1; j < n; j++ {
		f[0][j] += f[0][j-1] + j + 1
	}
	for i := 1; i < m; i++ {
		f[i][0] += f[i-1][0] + i + 1
		for j := 1; j < n; j++ {
			f[i][j] += min(f[i][j-1], f[i-1][j]) + (i+1)*(j+1)
		}
	}
	return int64(f[m-1][n-1])
}
