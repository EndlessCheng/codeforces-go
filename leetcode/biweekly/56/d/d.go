package main

// github.com/EndlessCheng/codeforces-go
func minCost(maxTime int, edges [][]int, fees []int) int {
	const inf int = 1e9
	n := len(fees)
	dp := make([][]int, maxTime+1)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = inf
		}
	}
	dp[0][0] = fees[0]
	// 原图是存在环的，如果直接在拆点图上跑最短路是需要用 Dijkstra 等最短路算法的
	// 但是注意到，若按照时间的升序转移，由于图中边权均为正，从当前时间出发是不可能转移到过去的时间上的，从而保证状态无后效性，也就无需使用最短路算法来求解了
	for t := 1; t <= maxTime; t++ {
		for _, e := range edges {
			if e[2] <= t {
				v, w, wt := e[0], e[1], e[2]
				dp[t][v] = min(dp[t][v], dp[t-wt][w]+fees[v])
				dp[t][w] = min(dp[t][w], dp[t-wt][v]+fees[w])
			}
		}
	}
	ans := inf
	for _, dv := range dp[1:] {
		ans = min(ans, dv[n-1])
	}
	if ans < inf {
		return ans
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
