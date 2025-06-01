package main

import (
	"math"
	"slices"
)

// https://space.bilibili.com/206214
func minAbsDiff(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	ans := make([][]int, m-k+1)
	arr := make([]int, k*k)
	for i := range ans {
		ans[i] = make([]int, n-k+1)
		for j := range ans[i] {
			a := arr[:0]
			for _, row := range grid[i : i+k] {
				a = append(a, row[j:j+k]...)
			}
			slices.Sort(a)

			res := math.MaxInt
			for p := 1; p < len(a); p++ {
				if a[p] > a[p-1] {
					res = min(res, a[p]-a[p-1])
				}
			}
			if res < math.MaxInt {
				ans[i][j] = res
			}
		}
	}
	return ans
}
