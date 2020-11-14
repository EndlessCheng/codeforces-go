package main

// github.com/EndlessCheng/codeforces-go
func minimumDeletions(s string) (ans int) {
	n := len(s)
	sum := make([]int, n+1)
	for i, b := range s {
		sum[i+1] = sum[i]
		if b == 'a' {
			sum[i+1]++
		}
	}
	ans = n
	for i, s := range sum {
		if del := sum[n] + i - 2*s; del < ans {
			ans = del
		}
	}
	return
}
