package main

// github.com/EndlessCheng/codeforces-go
func minSkips(dis []int, spd, hoursBefore int) int {
	n := len(dis)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 2e9
	}
	dp[0] = 0
	ceil := func(d int) int { return ((d-1)/spd + 1) * spd }
	for _, d := range dis {
		for j := n - 1; j > 0; j-- {
			dp[j] = min(ceil(dp[j]+d), dp[j-1]+d)
		}
		dp[0] += ceil(d)
	}
	for i, d := range dp {
		if (d-1)/spd < hoursBefore {
			return i
		}
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
