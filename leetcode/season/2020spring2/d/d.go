package main

func splitArray(a []int) (ans int) {
	const mx int = 1e6
	lpf := [mx + 1]int{1: 1}
	for i := 2; i <= mx; i++ {
		if lpf[i] == 0 {
			for j := i; j <= mx; j += i {
				if lpf[j] == 0 {
					lpf[j] = i
				}
			}
		}
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	n := len(a)
	dp := make([]int, n+1)
	mi := [mx]int{}
	for i := range mi {
		mi[i] = 1e9
	}
	for i, v := range a {
		dp[i+1] = dp[i] + 1
		for v > 1 {
			p := lpf[v]
			for v /= p; lpf[v] == p; v /= p {
			}
			dp[i+1] = min(dp[i+1], mi[p]+1)
			mi[p] = min(mi[p], dp[i])
		}
	}
	return dp[n]
}
