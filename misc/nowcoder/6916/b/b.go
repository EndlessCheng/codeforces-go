package main

// github.com/EndlessCheng/codeforces-go
func arrange(a, b []int) int {
	n := len(a)
	dp := make([][2]int, n)
	for i := range dp {
		dp[i] = [2]int{-1, -1}
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	var f func(int, int) int
	f = func(p, swap int) (res int) {
		if p == n-1 {
			return swap
		}
		dv := &dp[p][swap]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		res = 1e9
		if swap == 1 {
			a[p], b[p] = b[p], a[p]
		}
		if a[p] > a[p+1] && b[p] < b[p+1] {
			res = min(res, f(p+1, 0))
		}
		if a[p] > b[p+1] && b[p] < a[p+1] {
			res = min(res, f(p+1, 1))
		}
		if swap == 1 {
			a[p], b[p] = b[p], a[p]
			res++
		}
		return
	}
	ans := min(f(0, 0), f(0, 1))
	if ans >= 1e9 {
		ans = -1
	}
	return ans
}
