package main

// github.com/EndlessCheng/codeforces-go
func minimizeTheDifference(a [][]int, tar int) int {
	m := len(a)
	dp := make([]int, min(m*70, tar*2)+1)
	for i := range dp {
		dp[i] = -1e9
	}
	dp[0] = 0
	sumMin, sumMax := 0, 0
	for _, row := range a {
		mi, mx := row[0], row[0]
		for _, v := range row {
			if v > mx {
				mx = v
			} else if v < mi {
				mi = v
			}
		}
		sumMin += mi
		sumMax = min(sumMax+mx, tar*2)
		for j := sumMax; j > 0; j-- {
			for _, v := range row {
				if v <= j && dp[j-v]+1 > dp[j] {
					dp[j] = dp[j-v] + 1
				}
			}
		}
	}

	ans := abs(sumMin - tar)
	for i, v := range dp {
		if v == m {
			ans = min(ans, abs(i-tar))
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
