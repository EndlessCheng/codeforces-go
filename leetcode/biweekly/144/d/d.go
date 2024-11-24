package main

import "math"

// https://space.bilibili.com/206214
func maxCollectedFruits(fruits [][]int) (ans int) {
	n := len(fruits)
	f := make([][]int, n-1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	dp := func() int {
		for i := range f {
			for j := range f[i] {
				f[i][j] = math.MinInt
			}
		}
		f[0][n-1] = fruits[0][n-1]
		for i := 1; i < n-1; i++ {
			for j := max(n-1-i, i+1); j < n; j++ {
				f[i][j] = max(f[i-1][j-1], f[i-1][j], f[i-1][j+1]) + fruits[i][j]
			}
		}
		return f[n-2][n-1]
	}

	for i, row := range fruits {
		ans += row[i]
	}
	ans += dp()
	// 把下三角形中的数据填到上三角形中
	for i := range fruits {
		for j := range i {
			fruits[j][i] = fruits[i][j]
		}
	}
	return ans + dp()
}

func maxCollectedFruits2(fruits [][]int) (ans int) {
	n := len(fruits)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if j < n-1-i || j >= n {
			return math.MinInt
		}
		if i == 0 {
			return fruits[i][j]
		}
		p := &memo[i][j]
		if *p < 0 {
			*p = max(dfs(i-1, j-1), dfs(i-1, j), dfs(i-1, j+1)) + fruits[i][j]
		}
		return *p
	}

	for i, row := range fruits {
		ans += row[i]
	}

	ans += dfs(n-2, n-1) // 从下往上走，方便 1:1 翻译成递推

	// 把下三角形中的数据填到上三角形中
	for i := range fruits {
		for j := range i {
			fruits[j][i] = fruits[i][j]
		}
	}
	for i := range memo {
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	return ans + dfs(n-2, n-1)
}
