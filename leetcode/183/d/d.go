package main

func stoneGameIII(a []int) (ans string) {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n := len(a)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1e9
	}
	var f func(p int) int
	f = func(p int) (res int) {
		if p >= n {
			return 0
		}
		if dp[p] != 1e9 {
			return dp[p]
		}
		defer func() { dp[p] = res }()
		res = a[p] - f(p+1)
		if p+1 < n {
			res = max(res, a[p]+a[p+1]-f(p+2))
		}
		if p+2 < n {
			res = max(res, a[p]+a[p+1]+a[p+2]-f(p+3))
		}
		return
	}
	res := f(0)
	if res > 0 {
		return "Alice"
	}
	if res == 0 {
		return "Tie"
	}
	return "Bob"
}
