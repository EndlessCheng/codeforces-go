package main

import "sort"

// github.com/EndlessCheng/codeforces-go
var perm3 = [][]int{{0, 1, 2}, {0, 2, 1}, {1, 0, 2}, {1, 2, 0}, {2, 0, 1}, {2, 1, 0}}

func maxHeight(a [][]int) (ans int) {
	for _, v := range a {
		sort.Ints(v)
	}
	sort.Slice(a, func(i, j int) bool { a, b := a[i], a[j]; return a[0] < b[0] || a[0] == b[0] && (a[1] < b[1] || a[1] == b[1] && a[2] < b[2]) })
	b := make([][]int, 0, 6*len(a))
	for _, v := range a {
		for _, p := range perm3 {
			b = append(b, []int{v[p[0]], v[p[1]], v[p[2]]})
		}
	}
	dp := make([]int, len(b))
	for i, v := range b {
		for j, w := range b[:i-i%6] {
			if w[0] <= v[0] && w[1] <= v[1] && w[2] <= v[2] {
				dp[i] = max(dp[i], dp[j])
			}
		}
		dp[i] += v[0]
		ans = max(ans, dp[i])
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
