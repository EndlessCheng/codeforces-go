package main

import "slices"

// https://space.bilibili.com/206214
func subsequencePairCount(nums []int) int {
	const mod = 1_000_000_007
	n := len(nums)
	m := slices.Max(nums)
	f := [2][][]int{}
	for i := range f {
		f[i] = make([][]int, m+1)
		for j := range f[i] {
			f[i][j] = make([]int, m+1)
		}
	}
	for j := 1; j <= m; j++ {
		f[0][j][j] = 1
	}
	for i, x := range nums {
		for j := 0; j <= m; j++ {
			for k := 0; k <= m; k++ {
				f[(i+1)%2][j][k] = (f[i%2][j][k] + f[i%2][gcd(j, x)][k] + f[i%2][j][gcd(k, x)]) % mod
			}
		}
	}
	return f[n%2][0][0]
}

func subsequencePairCount(nums []int) int {
	const mod = 1_000_000_007
	n := len(nums)
	m := slices.Max(nums)
	f := make([][][]int, n+1)
	for i := range f {
		f[i] = make([][]int, m+1)
		for j := range f[i] {
			f[i][j] = make([]int, m+1)
		}
	}
	for j := 1; j <= m; j++ {
		f[0][j][j] = 1
	}
	for i, x := range nums {
		for j := 0; j <= m; j++ {
			for k := 0; k <= m; k++ {
				f[i+1][j][k] = (f[i][j][k] + f[i][gcd(j, x)][k] + f[i][j][gcd(k, x)]) % mod
			}
		}
	}
	return f[n][0][0]
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func subsequencePairCount3(nums []int) int {
	const mod = 1_000_000_007
	n := len(nums)
	m := slices.Max(nums)
	f := make([][][]int, n+1)
	for i := range f {
		f[i] = make([][]int, m+1)
		for j := range f[i] {
			f[i][j] = make([]int, m+1)
		}
	}
	for j := 1; j <= m; j++ {
		f[0][j][j] = 1
	}
	for i, x := range nums {
		for j := m; j >= 0; j-- {
			for k := m; k >= 0; k-- {
				f[i+1][j][k] = (f[i][j][k] + f[i][gcd(j, x)][k] + f[i][j][gcd(k, x)]) % mod
			}
		}
	}
	return f[n][0][0]
}

func subsequencePairCount2(nums []int) int {
	const mod = 1_000_000_007
	n := len(nums)
	m := slices.Max(nums)
	memo := make([][][]int, n)
	for i := range memo {
		memo[i] = make([][]int, m+1)
		for j := range memo[i] {
			memo[i][j] = make([]int, m+1)
			for k := range memo[i][j] {
				memo[i][j][k] = -1
			}
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, j, k int) int {
		if i < 0 {
			if j == k {
				return 1
			}
			return 0
		}
		p := &memo[i][j][k]
		if *p < 0 {
			*p = (dfs(i-1, j, k) + dfs(i-1, gcd(j, nums[i]), k) + dfs(i-1, j, gcd(k, nums[i]))) % mod
		}
		return *p
	}
	// 减去两个子序列都是空的情况
	return (dfs(n-1, 0, 0) - 1 + mod) % mod
}
