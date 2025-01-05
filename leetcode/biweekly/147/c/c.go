package main

import (
	"math"
	"slices"
)

// https://space.bilibili.com/206214
func longestSubsequence(nums []int) (ans int) {
	mx := slices.Max(nums)
	maxD := mx - slices.Min(nums)
	f := make([][]int, mx+1)
	for i := range f {
		f[i] = make([]int, maxD+1)
		for j := range f[i] {
			f[i][j] = math.MinInt
		}
	}

	for _, x := range nums {
		fx := 1
		for j := maxD; j >= 0; j-- {
			if x-j >= 0 {
				fx = max(fx, f[x-j][j]+1)
			}
			if x+j <= mx {
				fx = max(fx, f[x+j][j]+1)
			}
			f[x][j] = fx
			ans = max(ans, fx)
		}
	}
	return
}

func longestSubsequence2(nums []int) (ans int) {
	mx := slices.Max(nums)
	maxD := mx - slices.Min(nums)
	f := make([][]int, len(nums))
	for i := range f {
		f[i] = make([]int, maxD+2)
	}
	last := make([]int, mx+1)
	for i := range last {
		last[i] = -1
	}

	for i, x := range nums {
		for j := maxD; j >= 0; j-- {
			f[i][j] = max(f[i][j+1], 1)
			if x-j >= 0 && last[x-j] >= 0 {
				f[i][j] = max(f[i][j], f[last[x-j]][j]+1)
			}
			if x+j <= mx && last[x+j] >= 0 {
				f[i][j] = max(f[i][j], f[last[x+j]][j]+1)
			}
			ans = max(ans, f[i][j])
		}
		last[x] = i
	}
	return
}
