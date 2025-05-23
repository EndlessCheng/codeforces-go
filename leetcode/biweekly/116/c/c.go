package main

import "math"

// https://space.bilibili.com/206214
func lengthOfLongestSubsequence1(nums []int, target int) int {
	n := len(nums)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, target+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		if i < 0 {
			if j == 0 {
				return 0
			}
			return math.MinInt
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()
		if j < nums[i] {
			return dfs(i-1, j)
		}
		return max(dfs(i-1, j), dfs(i-1, j-nums[i])+1)
	}

	ans := dfs(n-1, target)
	if ans > 0 {
		return ans
	}
	return -1
}

func lengthOfLongestSubsequence2(nums []int, target int) int {
	n := len(nums)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, target+1)
	}
	for j := 1; j <= target; j++ {
		f[0][j] = math.MinInt
	}

	for i, x := range nums {
		for j := range target + 1 {
			if j < x {
				f[i+1][j] = f[i][j]
			} else {
				f[i+1][j] = max(f[i][j], f[i][j-x]+1)
			}
		}
	}

	if f[n][target] > 0 {
		return f[n][target]
	}
	return -1
}

func lengthOfLongestSubsequence3(nums []int, target int) int {
	f := make([]int, target+1)
	for i := 1; i <= target; i++ {
		f[i] = math.MinInt
	}
	for _, x := range nums {
		for j := target; j >= x; j-- {
			f[j] = max(f[j], f[j-x]+1)
		}
	}
	if f[target] > 0 {
		return f[target]
	}
	return -1
}

func lengthOfLongestSubsequence(nums []int, target int) int {
	f := make([]int, target+1)
	for i := 1; i <= target; i++ {
		f[i] = math.MinInt
	}
	s := 0
	for _, x := range nums {
		s = min(s+x, target)
		for j := s; j >= x; j-- {
			f[j] = max(f[j], f[j-x]+1)
		}
	}
	if f[target] > 0 {
		return f[target]
	}
	return -1
}
