package main

// github.com/EndlessCheng/codeforces-go
func getSumAbsoluteDifferences(a []int) (ans []int) {
	n := len(a)
	sum := make([]int, n+1)
	for i, v := range a {
		sum[i+1] = sum[i] + v
	}
	ans = make([]int, n)
	for i, v := range a {
		ans[i] = sum[n] - sum[i] - (n-i)*v + i*v - sum[i]
	}
	return
}
