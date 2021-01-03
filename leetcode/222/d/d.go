package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func minOperations(target, a []int) (ans int) {
	pos := map[int]int{}
	for i, v := range target {
		pos[v] = i + 1
	}
	dp := []int{}
	for _, v := range a {
		v = pos[v]
		if v == 0 {
			continue
		}
		if p := sort.SearchInts(dp, v); p < len(dp) {
			dp[p] = v
		} else {
			dp = append(dp, v)
		}
	}
	return len(target) - len(dp)
}
