package main

// github.com/EndlessCheng/codeforces-go
func dieSimulator(n int, rollMax []int) (ans int) {
	const mod int = 1e9 + 7
	dp := make([][6][]int, n)
	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = make([]int, rollMax[j])
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}
	var f func(int, int, int) int
	f = func(i, last, left int) (res int) {
		if i == n {
			return 1
		}
		dv := &dp[i][last][left]
		if *dv != -1 {
			return *dv
		}
		for j, mx := range rollMax {
			if j != last {
				res += f(i+1, j, mx-1)
			} else if left > 0 {
				res += f(i+1, j, left-1)
			}
		}
		res %= mod
		*dv = res
		return
	}
	for j, mx := range rollMax {
		ans += f(1, j, mx-1)
	}
	return ans % mod
}
