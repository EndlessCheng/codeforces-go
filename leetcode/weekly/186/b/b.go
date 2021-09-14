package main

func maxScore(a []int, k int) (ans int) {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n := len(a)
	sum := make([]int, n+1)
	for i, v := range a {
		sum[i+1] = sum[i] + v
	}
	for i := 0; i <= k; i++ {
		ans = max(ans, sum[i]+sum[n]-sum[n-k+i])
	}
	return
}
