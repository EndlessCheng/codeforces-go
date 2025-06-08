package main

import "math"

// https://space.bilibili.com/206214
func maximumProfit(prices []int, k int) int64 {
	f := make([][3]int, k+2)
	for j := 1; j <= k+1; j++ {
		f[j][1] = math.MinInt / 2
		f[j][2] = math.MinInt / 2
	}
	f[0][0] = math.MinInt / 2
	for _, p := range prices {
		for j := k + 1; j > 0; j-- {
			f[j][0] = max(f[j][0], f[j][1]+p, f[j][2]-p)
			f[j][1] = max(f[j][1], f[j-1][0]-p)
			f[j][2] = max(f[j][2], f[j-1][0]+p)
		}
	}
	return int64(f[k+1][0])
}

func maximumProfit2(prices []int, k int) int64 {
	n := len(prices)
	f := make([][][3]int, n+1)
	for i := range f {
		f[i] = make([][3]int, k+2)
		for j := range f[i] {
			f[i][j] = [3]int{math.MinInt / 2, math.MinInt / 2, math.MinInt / 2} // 防止溢出
		}
	}
	for j := 1; j <= k+1; j++ {
		f[0][j][0] = 0
	}
	for i, p := range prices {
		for j := 1; j <= k+1; j++ {
			f[i+1][j][0] = max(f[i][j][0], f[i][j][1]+p, f[i][j][2]-p)
			f[i+1][j][1] = max(f[i][j][1], f[i][j-1][0]-p)
			f[i+1][j][2] = max(f[i][j][2], f[i][j-1][0]+p)
		}
	}
	return int64(f[n][k+1][0])
}

func maximumProfit1(prices []int, k int) int64 {
	n := len(prices)
	memo := make([][][3]int, n)
	for i := range memo {
		memo[i] = make([][3]int, k+1)
		for j := range memo[i] {
			memo[i][j] = [3]int{-1, -1, -1} // -1 表示还没有计算过
		}
	}
	// 在 [0,i] 中完成至多 j 笔交易，第 i 天结束时的状态为 endState 的情况下的最大收益
	var dfs func(int, int, int) int
	dfs = func(i, j, endState int) (res int) {
		if j < 0 {
			return math.MinInt / 2 // 防止溢出
		}
		if i < 0 {
			if endState == 1 {
				return math.MinInt / 2
			}
			return
		}
		ptr := &memo[i][j][endState]
		if *ptr != -1 { // 之前计算过
			return *ptr
		}
		defer func() { *ptr = res }() // 记忆化
		p := prices[i]
		if endState == 0 {
			return max(dfs(i-1, j, 0), dfs(i-1, j, 1)+p, dfs(i-1, j, 2)-p)
		}
		if endState == 1 {
			return max(dfs(i-1, j, 1), dfs(i-1, j-1, 0)-p)
		}
		return max(dfs(i-1, j, 2), dfs(i-1, j-1, 0)+p)
	}
	return int64(dfs(n-1, k, 0))
}
