package main

func minDifficulty(jobDifficulty []int, d int) (ans int) {
	n := len(jobDifficulty)
	if n < d {
		return -1
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	const mx1 = 305
	const mx2 = 15
	dp := [mx1][mx1][mx2]int{}
	vis := [mx1][mx1][mx2]bool{}
	var f func(l, l0, curMax, left int) int
	f = func(l, l0, curMax, left int) (ans int) {
		if l == n && left == 0 {
			return 0
		}
		if l == n || left == 0 {
			return 1e9
		}
		if vis[l][l0][left] {
			return dp[l][l0][left]
		}
		vis[l][l0][left] = true
		defer func() { dp[l][l0][left] = ans }()
		curMax = max(curMax, jobDifficulty[l])
		f1 := f(l+1, l0, curMax, left)
		f2 := f(l+1, l+1, 0, left-1) + curMax
		return min(f1, f2)
	}
	return f(0, 0, 0, d)
}
