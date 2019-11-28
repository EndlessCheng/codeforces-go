package main

func numWays(steps int, arrLen int) int {
	const mx int = 505
	const mod int = 1e9 + 7

	dp := [mx][mx]int{}
	vis := [mx][mx]bool{}

	var f func(pos, left int) int
	f = func(pos, left int) int {
		if left == 0 {
			if pos == 0 {
				return 1
			}
			return 0
		}
		if pos < 0 || pos >= arrLen {
			return 0
		}
		if vis[pos][left] {
			return dp[pos][left]
		}

		ans := f(pos-1, left-1) + f(pos, left-1) + f(pos+1, left-1)
		dp[pos][left] = ans % mod
		vis[pos][left] = true
		return dp[pos][left]
	}
	return f(0, steps)
}
