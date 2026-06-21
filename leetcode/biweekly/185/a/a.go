package main

import "strings"

// https://space.bilibili.com/206214
func createGrid(m, n int) []string {
	ans := make([]string, m)
	ans[0] = strings.Repeat(".", n)
	for i := 1; i < m; i++ {
		ans[i] = strings.Repeat("#", n-1) + "."
	}
	return ans
}
