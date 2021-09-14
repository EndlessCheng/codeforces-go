package main

// github.com/EndlessCheng/codeforces-go
func getMaximumGenerated(n int) (ans int) {
	if n == 0 {
		return
	}
	a := make([]int, n+1)
	a[1] = 1
	for i := 2; i <= n; i++ {
		if i&1 == 0 {
			a[i] = a[i/2]
		} else {
			a[i] = a[i/2] + a[i/2+1]
		}
	}
	for _, v := range a {
		if v > ans {
			ans = v
		}
	}
	return
}
