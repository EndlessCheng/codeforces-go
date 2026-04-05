package main

import "math"

// https://space.bilibili.com/206214

// 非环形版本
func solve(a []int, k int) int {
	n := len(a)
	f := make([][]int, k+1)
	for i := range f {
		f[i] = make([]int, n)
		if i > 0 {
			f[i][0] = math.MaxInt / 2
			f[i][1] = math.MaxInt / 2
		}
	}

	for left := 1; left <= k; left++ {
		for i := 1; i < n-1; i++ {
			// 选或不选
			notChoose := f[left][i]
			choose := f[left-1][i-1] + max(max(a[i-1], a[i+1])-a[i]+1, 0)
			f[left][i+1] = min(notChoose, choose)
		}
	}

	return f[k][n-1]
}

func minOperations(nums []int, k int) int {
	n := len(nums)
	if k > n/2 {
		return -1
	}

	cnt := 0
	for i, x := range nums {
		if nums[(i-1+n)%n] < x && x > nums[(i+1)%n] {
			cnt++
		}
	}
	if cnt >= k { // 优化：已经有至少 k 个峰值了，无需操作
		return 0
	}

	// 如果 nums[0] 是峰顶，那么 nums[n-1] 不是峰顶
	ans1 := solve(append([]int{nums[n-1]}, nums...), k)
	// 如果 nums[0] 不是峰顶
	ans2 := solve(append(nums, nums[0]), k)
	return min(ans1, ans2)
}

//

// 非环形版本
func solve(a []int, k int) int {
	n := len(a)
	memo := make([][]int, k+1)
	for i := range memo {
		memo[i] = make([]int, n-1)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有计算过
		}
	}

	// 返回使 [0,i+1] 包含 left 个峰值的最小操作次数
	var dfs func(int, int) int
	dfs = func(left, i int) int {
		if left == 0 {
			return 0
		}
		if left > (i+1)/2 { // [0,i+1] 至多有 (i+1)/2 个峰值
			return math.MaxInt / 2 // 防止加法溢出
		}

		p := &memo[left][i]
		if *p != -1 { // 之前计算过
			return *p
		}

		// 选或不选
		notChoose := dfs(left, i-1)
		choose := dfs(left-1, i-2) + max(max(a[i-1], a[i+1])-a[i]+1, 0)
		res := min(notChoose, choose)

		*p = res // 记忆化
		return res
	}

	return dfs(k, n-2)
}

func minOperations(nums []int, k int) int {
	n := len(nums)
	if k > n/2 {
		return -1
	}

	// 如果 nums[0] 是峰顶，那么 nums[n-1] 不是峰顶
	ans1 := solve(append([]int{nums[n-1]}, nums...), k)
	// 如果 nums[0] 不是峰顶
	ans2 := solve(append(nums, nums[0]), k)
	return min(ans1, ans2)
}
