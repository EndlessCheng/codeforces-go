package main

// github.com/EndlessCheng/codeforces-go
func maxSum(a, b []int) (ans int) {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	c := [2][]int{a, b}
	pos := [2]map[int]int{{}, {}}
	for i, a := range c {
		for j, v := range a {
			pos[i][v] = j
		}
	}
	dp := make([][2]int, max(len(a), len(b)))
	var f func(int, int) int
	f = func(p, q int) (res int) {
		if p == len(c[q]) {
			return
		}
		dv := &dp[p][q]
		if *dv != 0 {
			return *dv
		}
		defer func() { *dv = res }()
		v := c[q][p]
		res = f(p+1, q)
		if i, ok := pos[q^1][v]; ok {
			res = max(res, f(i+1, q^1))
		}
		res += v
		return
	}
	return max(f(0, 0), f(0, 1)) % (1e9 + 7)
}
