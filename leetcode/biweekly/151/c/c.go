package main

import "slices"

// https://space.bilibili.com/206214
func minCost(nums []int) int {
	n := len(nums)
	var f []int
	if n%2 == 1 {
		f = slices.Clone(nums)
	} else {
		f = make([]int, n)
		for i, x := range nums {
			f[i] = max(x, nums[n-1])
		}
	}
	for i := n - 3 + n%2; i > 0; i -= 2 {
		b, c := nums[i], nums[i+1]
		for j, a := range nums[:i] {
			f[j] = min(f[j]+max(b, c), f[i]+max(a, c), f[i+1]+max(a, b))
		}
	}
	return f[0]
}

func minCost2(nums []int) int {
	n := len(nums)
	f := make([][]int, n+1)
	f[n] = nums
	f[n-1] = make([]int, n)
	for i, x := range nums {
		f[n-1][i] = max(x, nums[n-1])
	}
	for i := n - 3 + n%2; i > 0; i -= 2 {
		f[i] = make([]int, i)
		b, c := nums[i], nums[i+1]
		for j, a := range nums[:i] {
			f[i][j] = min(f[i+2][j]+max(b, c), f[i+2][i]+max(a, c), f[i+2][i+1]+max(a, b))
		}
	}
	return f[1][0]
}

func minCost1(nums []int) int {
	n := len(nums)
	memo := make([][]int, n-1)
	for i := range memo {
		memo[i] = make([]int, i)
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i == n {
			return nums[j]
		}
		if i == n-1 {
			return max(nums[j], nums[i])
		}
		p := &memo[i][j]
		a, b, c := nums[j], nums[i], nums[i+1]
		if *p == 0 {
			*p = min(dfs(i+2, j)+max(b, c), dfs(i+2, i)+max(a, c), dfs(i+2, i+1)+max(a, b))
		}
		return *p
	}
	return dfs(1, 0)
}
