package main

import "math"

// https://space.bilibili.com/206214
func maxAlternatingSum1(nums []int) int64 {
	n := len(nums)
	memo := make([][2][2]int, n)
	for i := range memo {
		memo[i] = [2][2]int{{math.MinInt, math.MinInt}, {math.MinInt, math.MinInt}}
	}

	// 只需在 1186 的基础上增加参数 rev
	var dfs func(int, int, int) int
	dfs = func(i, j, rev int) (res int) {
		if i == n {
			return math.MinInt / 2 // 除 2 防止负数相加溢出
		}
		p := &memo[i][j][rev]
		if *p != math.MinInt { // 之前计算过
			return *p
		}
		defer func() { *p = res }() // 记忆化

		x := nums[i]
		if rev > 0 {
			x = -x
		}

		if j == 0 {
			return max(dfs(i+1, 0, rev^1), 0) + x
		}
		return max(dfs(i+1, 1, rev^1)+x, dfs(i+1, 0, rev))
	}

	ans := math.MinInt
	for i := range nums {
		ans = max(ans, dfs(i, 0, 0), dfs(i, 1, 0))
	}
	return int64(ans)
}

func maxAlternatingSum2(nums []int) int64 {
	const negInf = math.MinInt / 2
	n := len(nums)
	f := make([][2][2]int, n+1)
	f[n] = [2][2]int{{negInf, negInf}, {negInf, negInf}} // 除 2 防止负数相加溢出
	ans := negInf
	for i := n - 1; i >= 0; i-- {
		x := nums[i]
		f[i][0][0] = max(f[i+1][0][1], 0) + x
		f[i][0][1] = max(f[i+1][0][0], 0) - x
		f[i][1][0] = max(f[i+1][1][1]+x, f[i+1][0][0])
		f[i][1][1] = max(f[i+1][1][0]-x, f[i+1][0][1])
		ans = max(ans, f[i][0][0], f[i][1][0])
	}
	return int64(ans)
}

func maxAlternatingSum(nums []int) int64 {
	const negInf = math.MinInt / 2
	f00, f01, f10, f11 := negInf, negInf, negInf, negInf
	ans := negInf
	for i := len(nums) - 1; i >= 0; i-- {
		x := nums[i]
		f10, f11 = max(f11+x, f00), max(f10-x, f01)
		f00, f01 = max(f01, 0)+x, max(f00, 0)-x
		ans = max(ans, f00, f10)
	}
	return int64(ans)
}
