package main

import (
	"math"
	"math/bits"
	"slices"
)

// https://space.bilibili.com/206214
func minRemovals1(nums []int, target int) int {
	m := bits.Len(uint(slices.Max(nums)))
	if m < bits.Len(uint(target)) {
		return -1
	}

	n := len(nums)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, 1<<m)
		for j := range f[i] {
			f[i][j] = math.MinInt
		}
	}
	f[0][0] = 0

	for i, x := range nums {
		for j := range 1 << m {
			f[i+1][j] = max(f[i][j], f[i][j^x]+1) // x 不选 or 选
		}
	}

	if f[n][target] < 0 {
		return -1
	}
	return len(nums) - f[n][target]
}

func minRemovals(nums []int, target int) int {
	m := bits.Len(uint(slices.Max(nums)))
	if m < bits.Len(uint(target)) {
		return -1
	}

	f := make([]int, 1<<m)
	for i := range f {
		f[i] = math.MinInt
	}
	f[0] = 0

	for _, x := range nums {
		nf := slices.Clone(f)
		for j := range 1 << m {
			nf[j] = max(nf[j], f[j^x]+1) // x 不选 or 选
		}
		f = nf
	}

	if f[target] < 0 {
		return -1
	}
	return len(nums) - f[target]
}
