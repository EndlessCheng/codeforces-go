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
	var f func(p int) int
	f = func(p int) (_ans int) {
		if dp[p] > 0 {
			return dp[p]
		}
		defer func() { dp[p] = _ans }()
		for i := p + 1; i <= p+d; i++ {
			if i == n || a[i] >= a[p] {
				break
			}
			_ans = max(_ans, f(i))
		}
		for i := p - 1; i >= p-d; i-- {
			if i == -1 || a[i] >= a[p] {
				break
			}
			_ans = max(_ans, f(i))
		}
		return _ans + 1
	}
	for i := range a {
		ans = max(ans, f(i))
	}
	return
}
