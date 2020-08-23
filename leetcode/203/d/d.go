package main

// github.com/EndlessCheng/codeforces-go
func stoneGameV(a []int) (ans int) {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n := len(a)
	sum := make([]int, n+1)
	for i, v := range a {
		sum[i+1] = sum[i] + v
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var f func(int, int) int
	f = func(l, r int) (res int) {
		if r-l == 1 {
			return
		}
		dv := &dp[l][r]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		for i := l + 1; i < r; i++ {
			sl, sr := sum[i]-sum[l], sum[r]-sum[i]
			var s int
			if sl == sr {
				s = sl + max(f(l, i), f(i, r))
			} else if sl < sr {
				s = sl + f(l, i)
			} else {
				s = sr + f(i, r)
			}
			res = max(res, s)
		}
		return
	}
	cnt:=0
	for i, row := range dp {
		for j, v := range row[i+1:] {
			_, _ = i, j
			if v == -1 {
				cnt++
			}
		}
	}
	println(cnt)
	return f(0, n)
}
