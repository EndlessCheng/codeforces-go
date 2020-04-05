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
	f = func(p int) (_ans int) {
		if p >= n {
			return 0
		}
		if dp[p] < 1e9 {
			return dp[p]
		}
		defer func() { dp[p] = _ans }()
		_ans = a[p] - f(p+1)
		if p+1 < n {
			_ans = max(_ans, a[p]+a[p+1]-f(p+2))
		}
		if p+2 < n {
			_ans = max(_ans, a[p]+a[p+1]+a[p+2]-f(p+3))
		}
		return
	}
	ret := f(0)
	if ret > 0 {
		return "Alice"
	}
	if ret == 0 {
		return "Tie"
	}
	return "Bob"
}
