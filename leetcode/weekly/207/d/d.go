package main

// github.com/EndlessCheng/codeforces-go
func connectTwoGroups(cost [][]int) (ans int) {
	n, m := len(cost), len(cost[0])
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	mi := append([]int(nil), cost[0]...)
	for _, row := range cost[1:] {
		for j, v := range row {
			mi[j] = min(mi[j], v)
		}
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 1<<m)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var f func(int, int) int
	f = func(p, conn int) (res int) {
		if p == n {
			for i, v := range mi {
				if conn>>i&1 == 0 {
					res += v
				}
			}
			return
		}
		dv := &dp[p][conn]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		res = 1e9
		for i, v := range cost[p] {
			res = min(res, v+f(p+1, conn|1<<i))
		}
		return
	}
	ans = f(0, 0)
	return
}
