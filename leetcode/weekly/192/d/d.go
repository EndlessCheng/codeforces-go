package main

func minCost(houses []int, cost [][]int, n int, m int, target int) (ans int) {
	const inf int = 1e9
	ans = inf
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	dp := make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, n+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, m)
			for k := range dp[i][j] {
				dp[i][j][k] = inf
			}
		}
	}
	if c := houses[0]; c > 0 {
		dp[0][1][c-1] = 0
	} else {
		dp[0][1] = cost[0]
	}
	for i := 1; i < n; i++ {
		for b := 1; b <= n; b++ {
			for c, cst := range cost[i] {
				if houses[i] > 0 {
					if c != houses[i]-1 {
						continue
					}
					cst = 0
				}
				for pc := 0; pc < m; pc++ {
					bb := b
					if c != pc {
						bb--
					}
					dp[i][b][c] = min(dp[i][b][c], dp[i-1][bb][pc]+cst)
				}
			}
		}
	}
	for _, v := range dp[n-1][target] {
		ans = min(ans, v)
	}
	if ans == inf {
		ans = -1
	}
	return
}
