package main

import "math"

// https://space.bilibili.com/206214

// 非环形版本
func solve1(a []int, k int) int {
	n := len(a)
	f := make([][]int, k+1)
	for i := range f {
		f[i] = make([]int, n)
	}
	for left := 1; left <= k; left++ {
		f[left][left*2-1] = math.MaxInt / 2
		for i := left*2 - 1; i < n-1-(k-left)*2; i++ {
			// 选或不选
			notChoose := f[left][i]
			choose := f[left-1][i-1] + max(max(a[i-1], a[i+1])-a[i]+1, 0)
			f[left][i+1] = min(notChoose, choose)
		}
	}
	return f[k][n-1]
}

// 非环形版本
func solve(a []int, k int) int {
	n := len(a)
	f := make([]int, n)
	for left := 1; left <= k; left++ {
		f0, f1 := f[left*2-2], f[left*2-1]
		f[left*2-1] = math.MaxInt / 2
		for i := left*2 - 1; i < n-1-(k-left)*2; i++ {
			// 选或不选
			notChoose := f[i]
			choose := f0 + max(max(a[i-1], a[i+1])-a[i]+1, 0)
			f0 = f1
			f1 = f[i+1]
			f[i+1] = min(notChoose, choose)
		}
	}
	return f[n-1]
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

	// 如果 nums[0] 是峰值，那么 nums[n-1] 不是峰值
	ans1 := solve(append([]int{nums[n-1]}, nums...), k)
	// 如果 nums[0] 不是峰值
	ans2 := solve(append(nums, nums[0]), k)
	return min(ans1, ans2)
}
