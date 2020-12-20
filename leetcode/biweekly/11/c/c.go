package main

// github.com/EndlessCheng/codeforces-go
func probabilityOfHeads(prob []float64, target int) (ans float64) {
	n := len(prob)
	dp := make([][]float64, n)
	for i := range dp {
		dp[i] = make([]float64, target+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var f func(int, int) float64
	f = func(p, left int) (res float64) {
		if left < 0 || left > n-p {
			return
		}
		if p == n {
			return 1
		}
		dv := &dp[p][left]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		return prob[p]*f(p+1, left-1) + (1-prob[p])*f(p+1, left)
	}
	ans = f(0, target)
	return
}
