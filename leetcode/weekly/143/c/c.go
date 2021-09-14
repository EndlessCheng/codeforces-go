package main

// github.com/EndlessCheng/codeforces-go
func minHeightShelves(a [][]int, w int) (ans int) {
	n := len(a)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = -1
	}
	var f func(int) int
	f = func(p int) (res int) {
		if p == n {
			return
		}
		dv := &dp[p]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		res = 1e9
		for i, mx, s := p, 0, 0; i < n && s+a[i][0] <= w; i++ {
			mx = max(mx, a[i][1])
			s += a[i][0]
			res = min(res, mx+f(i+1))
		}
		return
	}
	ans = f(0)
	return
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
