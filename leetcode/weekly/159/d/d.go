package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func jobScheduling(startTime, endTime, profit []int) (ans int) {
	type job struct{ l, r, p int }
	a := make([]job, len(startTime))
	for i, l := range startTime {
		a[i] = job{l, endTime[i], profit[i]}
	}
	sort.Slice(a, func(i, j int) bool { return a[i].r < a[j].r })

	type pair struct{ p, v int }
	dp := []pair{{}}
	for _, p := range a {
		i := sort.Search(len(dp), func(i int) bool { return dp[i].p > p.l }) - 1
		if dp[i].v+p.p > dp[len(dp)-1].v {
			dp = append(dp, pair{p.r, dp[i].v + p.p})
		}
	}
	return dp[len(dp)-1].v
}
