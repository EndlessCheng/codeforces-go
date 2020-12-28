package main

// github.com/EndlessCheng/codeforces-go
func dietPlanPerformance(a []int, k, lower, upper int) (ans int) {
	n := len(a)
	sum := make([]int, n+1)
	for i, v := range a {
		sum[i+1] = sum[i] + v
	}
	for i := k; i <= n; i++ {
		if d := sum[i] - sum[i-k]; d < lower {
			ans--
		} else if d > upper {
			ans++
		}
	}
	return
}
