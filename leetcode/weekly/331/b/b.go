package main

import "strings"

// https://space.bilibili.com/206214
func vowelStrings(words []string, queries [][]int) []int {
	sum := make([]int, len(words)+1)
	for i, w := range words {
		sum[i+1] = sum[i]
		if strings.Contains("aeiou", w[:1]) && strings.Contains("aeiou", w[len(w)-1:]) {
			sum[i+1]++
		}
	}
	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = sum[q[1]+1] - sum[q[0]]
	}
	return ans
}
