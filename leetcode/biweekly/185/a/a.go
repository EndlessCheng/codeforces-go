package main

import "strings"

// https://space.bilibili.com/206214
func createGrid(m, n int) []string {
	ans := make([]string, m)
	ans[0] = strings.Repeat(".", n)
	row := strings.Repeat("#", n-1) + "." // 避免在循环中反复创建字符串
	for i := 1; i < m; i++ {
		ans[i] = row
	}
	return ans
}
