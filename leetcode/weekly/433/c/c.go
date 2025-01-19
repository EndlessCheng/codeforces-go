package main

import (
	"math"
	"slices"
)

func minCost(n int, cost [][]int) int64 {
	f := make([][3][3]int, n/2+1)
	for i, row := range cost[:n/2] {
		row2 := cost[n-1-i]
		for preJ := range 3 {
			for preK := range 3 {
				res := math.MaxInt
				for j, c1 := range row {
					if j == preJ {
						continue
					}
					for k, c2 := range row2 {
						if k != preK && k != j {
							res = min(res, f[i][j][k]+c1+c2)
						}
					}
				}
				f[i+1][preJ][preK] = res
			}
		}
	}

	// 枚举所有初始颜色，取最小值
	res := math.MaxInt
	for _, row := range f[n/2] {
		res = min(res, slices.Min(row[:]))
	}
	return int64(res)
}

func minCost2(n int, cost [][]int) int64 {
	memo := make([][4][4]int, n/2)
	for i := range memo {
		for j := range memo[i] {
			for k := range memo[i][j] {
				memo[i][j][k] = -1
			}
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, preJ, preK int) (res int) {
		if i < 0 {
			return
		}
		p := &memo[i][preJ][preK]
		if *p != -1 {
			return *p
		}
		res = math.MaxInt
		for j, c1 := range cost[i] {
			if j == preJ {
				continue
			}
			for k, c2 := range cost[n-1-i] {
				if k != preK && k != j {
					res = min(res, dfs(i-1, j, k)+c1+c2)
				}
			}
		}
		*p = res
		return
	}
	return int64(dfs(n/2-1, 3, 3))
}
