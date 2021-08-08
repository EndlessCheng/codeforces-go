package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func longestObstacleCourseAtEachPosition(obstacles []int) []int {
	ans := make([]int, len(obstacles))
	dp := []int{}
	for i, v := range obstacles {
		p := sort.SearchInts(dp, v+1)
		if p < len(dp) {
			dp[p] = v
		} else {
			dp = append(dp, v)
		}
		ans[i] = p + 1
	}
	return ans
}
