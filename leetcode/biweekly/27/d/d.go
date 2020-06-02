package main

func cherryPickup(g [][]int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n, m := len(g), len(g[0])
	dp := make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, m)
		for j := range dp[i] {
			dp[i][j] = make([]int, m)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}
	var f func(int, int, int) int
	f = func(x int, y1 int, y2 int) (res int) {
		if x == n || !(0 <= y1 && y1 < m) || !(0 <= y2 && y2 < m) {
			return
		}
		dv := &dp[x][y1][y2]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				res = max(res, f(x+1, y1+i, y2+j))
			}
		}
		res += g[x][y1]
		if y1 != y2 {
			res += g[x][y2]
		}
		return
	}
	return f(0, 0, m-1)
}
