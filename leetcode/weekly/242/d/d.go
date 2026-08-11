package main

import "math"

// github.com/EndlessCheng/codeforces-go
func stoneGameVIII1(stones []int) int {
	n := len(stones)
	sum := make([]int, n)
	sum[0] = stones[0]
	for i := 1; i < n; i++ {
		sum[i] = sum[i-1] + stones[i]
	}
	memo := make([]int, n-1)
	for i := range memo {
		memo[i] = math.MaxInt
	}

	var dfs func(int) int
	dfs = func(i int) int {
		if i == n-1 {
			return sum[n-1]
		}
		p := &memo[i]
		if *p == math.MaxInt {
			*p = max(dfs(i+1), sum[i]-dfs(i+1))
		}
		return *p
	}

	return dfs(1)
}

func stoneGameVIII2(stones []int) int {
	n := len(stones)
	sum := make([]int, n)
	sum[0] = stones[0]
	for i := 1; i < n; i++ {
		sum[i] = sum[i-1] + stones[i]
	}

	f := make([]int, n)
	f[n-1] = sum[n-1]
	for i := n - 2; i > 0; i-- {
		f[i] = max(f[i+1], sum[i]-f[i+1])
	}
	return f[1]
}

func stoneGameVIII(stones []int) int {
	sum := 0
	for _, x := range stones {
		sum += x
	}

	f := sum
	for i := len(stones) - 2; i > 0; i-- {
		sum -= stones[i+1]
		f = max(f, sum-f)
	}
	return f
}
