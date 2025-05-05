package main

import "math"

// https://space.bilibili.com/206214
func minTravelTime1(_, n, K int, position, time []int) int {
	s := make([]int, n)
	for i, t := range time[:n-1] { // time[n-1] 用不到
		s[i+1] = s[i] + t // 计算 time 的前缀和
	}

	memo := make([][][]int, n-1)
	for i := range memo {
		memo[i] = make([][]int, n-1)
		for j := range memo[i] {
			memo[i][j] = make([]int, K+1)
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, j, leftK int) int {
		if j == n-1 { // 到达终点
			if leftK > 0 { // 不合法
				return math.MaxInt / 2 // 避免下面计算 r 的地方加法溢出
			}
			return 0
		}
		p := &memo[i][j][leftK]
		if *p > 0 {
			return *p
		}
		res := math.MaxInt
		t := s[j+1] - s[i] // 合并到 time[i] 的时间
		// 枚举下一个子数组 [j+1, k]
		for k := j + 1; k < min(n, j+2+leftK); k++ {
			r := dfs(j+1, k, leftK-(k-j-1)) + (position[k]-position[j])*t
			res = min(res, r)
		}
		*p = res
		return res
	}
	return dfs(0, 0, K) // 第一个子数组是 [0, 0]
}

func minTravelTime(_, n, K int, position, time []int) int {
	s := make([]int, n)
	for i, t := range time[:n-1] { // time[n-1] 用不到
		s[i+1] = s[i] + t // 计算 time 的前缀和
	}

	f := make([][][]int, n)
	for i := range f {
		f[i] = make([][]int, n)
		for j := range f[i] {
			f[i][j] = make([]int, K+1)
		}
		for leftK := 1; leftK <= K; leftK++ {
			f[i][n-1][leftK] = math.MaxInt / 2
		}
	}
	for i := n - 2; i >= 0; i-- { // 转移来源（j+1）比 i 大，所以要倒序
		for j := i; j < n-1; j++ {
			t := s[j+1] - s[i] // 合并到 time[i] 的时间
			for leftK := range K + 1 {
				res := math.MaxInt
				// 枚举下一个子数组 [j+1, k]
				for k := j + 1; k < min(n, j+2+leftK); k++ {
					r := f[j+1][k][leftK-(k-j-1)] + (position[k]-position[j])*t
					res = min(res, r)
				}
				f[i][j][leftK] = res
			}
		}
	}
	return f[0][0][K] // 第一个子数组是 [0, 0]
}
