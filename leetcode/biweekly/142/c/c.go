package main

import "slices"

// https://space.bilibili.com/206214
func maxScore(n, k int, stayScore, travelScore [][]int) int {
	f := make([][]int, k+1)
	f[k] = make([]int, n)
	for i, row := range slices.Backward(stayScore) {
		f[i] = make([]int, n)
		for j, s := range row {
			f[i][j] = f[i+1][j] + s
			for d, ts := range travelScore[j] {
				f[i][j] = max(f[i][j], f[i+1][d]+ts)
			}
		}
	}
	return slices.Max(f[0])
}

func maxScore2(n, k int, stayScore, travelScore [][]int) (ans int) {
	memo := make([][]int, k)
	for i := range memo {
		memo[i] = make([]int, n)
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i == k {
			return 0
		}
		p := &memo[i][j]
		if *p > 0 {
			return *p
		}
		res := dfs(i+1, j) + stayScore[i][j] // 留在当前城市 j
		for d, s := range travelScore[j] {
			// 注意题目保证 travelScore[i][i] = 0，这一定不如留在当前城市优
			res = max(res, dfs(i+1, d)+s) // 前往另外一座城市 d
		}
		*p = res
		return res
	}
	for j := range n {
		ans = max(ans, dfs(0, j)) // 选择城市 j 作为起点
	}
	return
}
