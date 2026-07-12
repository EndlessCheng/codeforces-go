package main

import (
	"bytes"
	"strings"
)

// https://space.bilibili.com/206214
func createGrid(m, n, k int) []string {
	// 特判
	if k == 4 && m == 3 && n == 3 {
		return []string{"..#", "...", "#.."}
	}

	if m == 1 || n == 1 {
		// 一行或一列，只能有一种方案
		if k > 1 {
			return nil
		}
		ans := make([]string, m)
		row := strings.Repeat(".", n)
		for i := range ans {
			ans[i] = row
		}
		return ans
	}

	// 至少要有 k 行或 k 列（特殊情况上面已判断）
	if m < k && n < k {
		return nil
	}

	// 初始全为 '#'
	a := make([][]byte, m)
	for i := range a {
		a[i] = bytes.Repeat([]byte{'#'}, n)
	}
	if m >= k { // 至少有 k 行
		// 第一列改成 '.'
		for _, row := range a {
			row[0] = '.'
		}
		// 第二列末尾 k 个 '.'
		for _, row := range a[m-k : m] {
			row[1] = '.'
		}
		// 最后一行改成 '.'
		a[m-1] = bytes.Repeat([]byte{'.'}, n)
	} else { // 至少有 k 列
		// 第一行改成 '.'
		a[0] = bytes.Repeat([]byte{'.'}, n)
		// 第二行末尾 k 个 '.'
		for j := n - k; j < n; j++ {
			a[1][j] = '.'
		}
		// 最后一列改成 '.'
		for _, row := range a {
			row[n-1] = '.'
		}
	}

	ans := make([]string, m)
	for i, row := range a {
		ans[i] = string(row)
	}
	return ans
}
