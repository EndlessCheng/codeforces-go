package main

import "math"

func maxDotProduct1(nums1, nums2 []int) int {
	n := len(nums1)
	m := len(nums2)

	// memo[i][j] 表示从 nums1[:i+1] 和 nums2[:j+1] 中选两个长度相同的【非空】子序列的最大点积
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, m)
		for j := range memo[i] {
			memo[i][j] = math.MaxInt
		}
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || j < 0 {
			// 其中一个数组没有元素，无法选出非空子序列
			return math.MinInt // 下面计算 max 不会取到无解情况
		}

		p := &memo[i][j]
		if *p != math.MaxInt {
			return *p
		}

		// 选 nums1[i] 和 nums2[j]
		// 和前面的子序列拼起来，或者不拼（作为子序列的第一个数）
		res1 := max(dfs(i-1, j-1), 0) + nums1[i]*nums2[j]

		// 不选 nums1[i]
		res2 := dfs(i-1, j)

		// 不选 nums2[j]
		res3 := dfs(i, j-1)

		*p = max(res1, res2, res3)
		return *p
	}

	return dfs(n-1, m-1)
}

func maxDotProduct2(nums1, nums2 []int) int {
	n := len(nums1)
	m := len(nums2)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, m+1)
		for j := range f[i] {
			f[i][j] = math.MinInt
		}
	}

	for i, x := range nums1 {
		for j, y := range nums2 {
			f[i+1][j+1] = max(max(f[i][j], 0)+x*y, f[i][j+1], f[i+1][j])
		}
	}
	return f[n][m]
}

func maxDotProduct(nums1, nums2 []int) int {
	m := len(nums2)
	f := make([]int, m+1)
	for i := range f {
		f[i] = math.MinInt
	}
	for _, x := range nums1 {
		pre := f[0]
		for j, y := range nums2 {
			tmp := f[j+1]
			f[j+1] = max(max(pre, 0)+x*y, f[j+1], f[j])
			pre = tmp
		}
	}
	return f[m]
}
