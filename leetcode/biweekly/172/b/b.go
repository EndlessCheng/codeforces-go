package main

import "math"

// https://space.bilibili.com/206214
func maximumSum(nums []int) int {
	const K = 3
	const MOD = 3
	f := [K + 1][MOD]int{}
	for i := range f {
		for j := range f[i] {
			f[i][j] = math.MinInt
		}
	}
	f[0][0] = 0
	for _, x := range nums {
		for j := K; j > 0; j-- {
			for r := range MOD {
				f[j][r] = max(f[j][r], f[j-1][(r-x%MOD+MOD)%MOD]+x)
			}
		}
	}
	return max(f[K][0], 0)
}

func maximumSum2(nums []int) int {
	const K = 3
	const MOD = 3
	f := [K + 1][MOD]int{}
	for i := range f {
		for j := range f[i] {
			f[i][j] = math.MinInt
		}
	}
	f[0][0] = 0
	for _, x := range nums {
		for j := K - 1; j >= 0; j-- {
			for r := range MOD {
				f[j+1][(r+x)%MOD] = max(f[j+1][(r+x)%MOD], f[j][r]+x)
			}
		}
	}
	return max(f[K][0], 0)
}
