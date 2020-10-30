package main

// github.com/EndlessCheng/codeforces-go
func largestNumber(weights []int, tar int) string {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dp := make([]int, tar+1)
	for i := range dp {
		dp[i] = -1e9
	}
	dp[0] = 0
	for _, w := range weights {
		for j := w; j <= tar; j++ {
			dp[j] = max(dp[j], dp[j-w]+1)
		}
	}
	if dp[tar] < 0 {
		return "0"
	}

	ans := []byte{}
	for i := 8; i >= 0; i-- {
		for w := weights[i]; tar >= w && dp[tar-w]+1 == dp[tar]; tar -= w {
			ans = append(ans, byte('0'+i+1))
		}
	}
	return string(ans)
}
