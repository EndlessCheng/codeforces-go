package main

// github.com/EndlessCheng/codeforces-go
func numberWays(hats [][]int) int {
	man := [41][]int{}
	for i, a := range hats {
		for _, v := range a {
			man[v] = append(man[v], i)
		}
	}
	m := 1 << len(hats)
	dp := make([][]int, 41)
	for i := range dp {
		dp[i] = make([]int, m)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var f func(int, int) int
	f = func(p, set int) (res int) {
		if set == m-1 {
			return 1
		}
		if p == 41 {
			return
		}
		dv := &dp[p][set]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		res = f(p+1, set)
		for _, v := range man[p] {
			if set>>v&1 == 0 {
				res += f(p+1, set|1<<v)
			}
		}
		return res % (1e9 + 7)
	}
	return f(1, 0)
}
