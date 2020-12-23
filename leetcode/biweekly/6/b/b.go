package main

// github.com/EndlessCheng/codeforces-go
func minSwaps(a []int) (ans int) {
	n := len(a)
	sum := make([]int, n+1)
	for i, v := range a {
		sum[i+1] = sum[i] + v
	}
	cnt := sum[n]
	ans = n
	for i := cnt; i <= n; i++ {
		ans = min(ans, cnt-sum[i]+sum[i-cnt])
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
