package main

import (
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func gcdValues(nums []int, queries []int64) []int {
	mx := slices.Max(nums)
	cntX := make([]int, mx+1)
	for _, x := range nums {
		cntX[x]++
	}

	cntGcd := make([]int, mx+1)
	for i := mx; i > 0; i-- {
		c := 0
		for j := i; j <= mx; j += i {
			c += cntX[j]
			cntGcd[i] -= cntGcd[j] // gcd 是 2i,3i,4i,... 的数对不能统计进来
		}
		cntGcd[i] += c * (c - 1) / 2 // c 个数选 2 个，组成 c*(c-1)/2 个数对
	}

	for i := 2; i <= mx; i++ {
		cntGcd[i] += cntGcd[i-1] // 前缀和
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = sort.SearchInts(cntGcd, int(q)+1)
	}
	return ans
}
