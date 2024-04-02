package main

import (
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func sumOfPowers(nums []int, k int) (ans int) {
	const mod = 1_000_000_007
	calc := func(a []int, lowDiff int) []int {
		n := len(a)
		memo := make([][]int, n)
		for i := range memo {
			memo[i] = make([]int, n+1)
			for j := range memo[i] {
				memo[i][j] = -1
			}
		}
		var dfs func(int, int) int
		dfs = func(i, j int) (res int) {
			if j > i+1 {
				return
			}
			if j == 0 {
				return 1
			}
			p := &memo[i][j]
			if *p != -1 {
				return *p
			}
			for k, v := range a[:i] {
				if a[i]-v >= lowDiff {
					res += dfs(k, j-1)
				}
			}
			res %= mod
			*p = res
			return
		}
		res := make([]int, k-1)
		res[0] = 1
		for j := 1; j < k-1; j++ {
			for i := range a {
				res[j] += dfs(i, j-1)
			}
			res[j] %= mod
		}
		return res
	}

	slices.Sort(nums)
	for i, x := range nums {
		for _, y := range nums[i+1:] {
			d := y - x // 能量
			res1 := calc(nums[:sort.SearchInts(nums[:i], x-d+1)], d)
			res2 := calc(nums[sort.SearchInts(nums, y+d+1):], d+1)
			for l, res := range res1 {
				ans = (ans + d*res%mod*res2[k-2-l]) % mod
			}
		}
	}
	return
}
