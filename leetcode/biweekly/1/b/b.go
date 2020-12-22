package main

import (
	"index/suffixarray"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func indexPairs(text string, words []string) (ans [][]int) {
	sa := suffixarray.New([]byte(text))
	for _, w := range words {
		for _, p := range sa.Lookup([]byte(w), -1) {
			ans = append(ans, []int{p, p + len(w) - 1})
		}
	}
	sort.Slice(ans, func(i, j int) bool { a, b := ans[i], ans[j]; return a[0] < b[0] || a[0] == b[0] && a[1] < b[1] })
	return
}
