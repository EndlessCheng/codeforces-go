package main

func maxJumps(a []int, d int) (ans int) {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n := len(a)

	dp := make([]int, n)
	for i := range dp {
		dp[i] = -1
	}
	var f func(p int) int
	f = func(p int) (_ans int) {
		if dp[p] != -1 {
			return dp[p]
		}
		defer func() { dp[p] = _ans }()
		ap := a[p]
		hMax := 0
		for i := p + 1; i <= p+d; i++ {
			if i == n || a[i] >= ap {
				break
			}
			_ans = max(_ans, f(i))
			hMax = max(hMax, a[i])
		}
		hMax = 0
		for i := p - 1; i >= p-d; i-- {
			if i == -1 || a[i] >= ap {
				break
			}
			_ans = max(_ans, f(i))
			hMax = max(hMax, a[i])
		}
		return _ans + 1
	}
	for i := range a {
		ans = max(ans, f(i))
	}
	return
}
