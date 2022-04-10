package main

import "sort"

// github.com/EndlessCheng/codeforces-go
const inf int = 1e18

func goShopping(a []int, b []int) int {
	n := len(a)
	type pair struct{ pa, pb int }
	ps := make([]pair, n)
	for i, v := range a {
		ps[i] = pair{v, b[i]}
	}
	sort.Slice(ps, func(i, j int) bool { return ps[i].pb > ps[j].pb })

	ans := inf
	dp1 := newDP1()
	dp1[0][0] = 0
	for _, p := range ps {
		pa, pb := p.pa, p.pb
		ndp := newDP1()
		for j := range dp1 {
			for k := range dp1[j] {
				if dp1[j][k] < inf {
					ndp[min(j+1, 3)][k] = min(ndp[min(j+1, 3)][k], dp1[j][k]+pa*7)
					v := dp1[j][k]
					if k < 2 {
						v += pb * 10
					}
					ndp[j][(k+1)%3] = min(ndp[j][(k+1)%3], v)
				}
			}
		}
		dp1 = ndp
	}
	for _, v := range dp1[3] {
		ans = min(ans, v)
	}

	dp2 := newDP2()
	dp2[0][0] = 0
	for _, p := range ps {
		pa, pb := p.pa, p.pb
		ndp := newDP2()
		for j := range dp2 {
			for k := range dp2[j] {
				if dp2[j][k] < inf {
					if j < 2 {
						ndp[j+1][k] = min(ndp[j+1][k], dp2[j][k]+pa*10)
					}
					v := dp2[j][k]
					if k < 2 {
						v += pb * 10
					}
					ndp[j][(k+1)%3] = min(ndp[j][(k+1)%3], v)
				}
			}
		}
		dp2 = ndp
	}
	for _, r := range dp2 {
		for _, v := range r {
			ans = min(ans, v)
		}
	}
	return ans / 10
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func newDP1() (dp [4][3]int) {
	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = inf
		}
	}
	return
}

func newDP2() (dp [3][3]int) {
	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = inf
		}
	}
	return
}

