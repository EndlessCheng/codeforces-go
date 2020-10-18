package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func bestTeamScore(scores []int, ages []int) (ans int) {
	n := len(scores)
	type pair struct{ s, age int }
	ps := make([]pair, n)
	for i, s := range scores {
		ps[i] = pair{s, ages[i]}
	}
	sort.Slice(ps, func(i, j int) bool { a, b := ps[i], ps[j]; return a.s < b.s || a.s == b.s && a.age < b.age })

	dp := make([][1001]int, n)
	for i := range dp {
		for j := 0; j <= 1000; j++ {
			dp[i][j] = -1
		}
	}
	var f func(int, int) int
	f = func(i, maxAge int) (res int) {
		if i == n {
			return
		}
		dv := &dp[i][maxAge]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		res = f(i+1, maxAge)
		if p := ps[i]; p.age >= maxAge {
			res = max(res, p.s+f(i+1, p.age))
		}
		return
	}
	ans = f(0, 0)
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
