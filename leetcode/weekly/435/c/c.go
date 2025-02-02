package main

import "math"

// https://space.bilibili.com/206214
func minimumIncrements(nums []int, target []int) int {
	// 预处理 target 的所有子集的 LCM
	m := len(target)
	lcms := make([]int, 1<<m)
	lcms[0] = 1
	for i, t := range target {
		bit := 1 << i
		for mask, l := range lcms[:bit] {
			lcms[bit|mask] = lcm(t, l)
		}
	}

	f := make([]int, 1<<m)
	for j := 1; j < 1<<m; j++ {
		f[j] = math.MaxInt / 2
	}
	for _, x := range nums {
		for j := 1<<m - 1; j > 0; j-- {
			for sub := j; sub > 0; sub = (sub - 1) & j {
				l := lcms[sub]
				f[j] = min(f[j], f[j^sub]+(l-x%l)%l)
			}
		}
	}
	return f[1<<m-1]
}

func gcd(a, b int) int { for a != 0 { a, b = b%a, a }; return b }
func lcm(a, b int) int { return a / gcd(a, b) * b }

func minimumIncrements2(nums []int, target []int) int {
	// 预处理 target 的所有子集的 LCM
	m := len(target)
	lcms := make([]int, 1<<m)
	lcms[0] = 1
	for i, t := range target {
		bit := 1 << i
		for mask, l := range lcms[:bit] {
			lcms[bit|mask] = lcm(t, l)
		}
	}

	n := len(nums)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, 1<<m)
	}
	for j := 1; j < 1<<m; j++ {
		f[0][j] = math.MaxInt / 2
	}
	for i, x := range nums {
		for j := 1; j < 1<<m; j++ {
			// 不修改 nums[i]
			f[i+1][j] = f[i][j]
			// 枚举 j 的所有非空子集 sub，把 nums[i] 改成 lcms[sub] 的倍数
			for sub := j; sub > 0; sub = (sub - 1) & j {
				l := lcms[sub]
				f[i+1][j] = min(f[i+1][j], f[i][j^sub]+(l-x%l)%l)
			}
		}
	}
	return f[n][1<<m-1]
}

func minimumIncrements1(nums []int, target []int) int {
	// 计算 target 的所有子集的 LCM
	m := len(target)
	lcms := make([]int, 1<<m)
	lcms[0] = 1
	for i, t := range target {
		bit := 1 << i
		for mask, l := range lcms[:bit] {
			lcms[bit|mask] = lcm(t, l)
		}
	}

	n := len(nums)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, 1<<m)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		if j == 0 {
			return
		}
		if i < 0 { // 不能有剩余元素
			return math.MaxInt / 2
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()
		// 不修改 nums[i]
		res = dfs(i-1, j)
		// 枚举 j 的所有非空子集 sub，把 nums[i] 改成 lcms[sub] 的倍数
		for sub := j; sub > 0; sub = (sub - 1) & j {
			l := lcms[sub]
			res = min(res, dfs(i-1, j^sub)+(l-nums[i]%l)%l)
		}
		return
	}
	return dfs(n-1, 1<<m-1)
}
