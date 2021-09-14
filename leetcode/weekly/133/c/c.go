package main

// github.com/EndlessCheng/codeforces-go
func maxSumTwoNoOverlap(a []int, l, m int) (ans int) {
	n := len(a)
	s := make([]int, n+1)
	for i, v := range a {
		s[i+1] = s[i] + v
	}
	f := func(l, m int) {
		mx := make([]int, n)
		for i := l; i < n; i++ {
			mx[i] = max(mx[i-1], s[i]-s[i-l])
		}
		v := 0
		for i := n - m; i > 0; i-- {
			v = max(v, s[i+m]-s[i])
			ans = max(ans, v+mx[i])
		}
	}
	f(l, m)
	f(m, l)
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
