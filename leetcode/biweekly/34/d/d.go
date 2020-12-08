package main

// github.com/EndlessCheng/codeforces-go
func countRoutes(a []int, start, finish, m int) int {
	n := len(a)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var f func(int, int) int
	f = func(p, left int) (res int) {
		if left < 0 {
			return
		}
		dv := &dp[p][left]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		if p == finish {
			res = 1
		}
		for i, v := range a {
			if i != p {
				res += f(i, left-abs(a[p]-v))
			}
		}
		return res % (1e9 + 7)
	}
	return f(start, m)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
