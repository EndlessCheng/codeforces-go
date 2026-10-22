package main

import "math"

// https://space.bilibili.com/206214
func minOperations1(nums []int) int {
	n := len(nums)
	cnt := 0
	p := 0
	for i := 1; i < n && cnt < 2; i++ {
		if nums[i-1] > nums[i] {
			cnt++
			p = i
		}
	}

	if cnt == 0 { // 已是递增
		return 0
	}
	ans := math.MaxInt
	if cnt == 1 && nums[0] > nums[n-1] { // 两个递增段
		ans = min(p, n-p+2)
	}

	cnt = 0
	p = 0
	for i := 1; i < n && cnt < 2; i++ {
		if nums[i-1] < nums[i] {
			cnt++
			p = i
		}
	}

	if cnt == 0 { // 已是递减
		return 1
	}
	if cnt == 1 && nums[0] < nums[n-1] { // 两个递减段
		ans = min(ans, p+1, n-p+1)
	}

	if ans == math.MaxInt {
		return -1
	}
	return ans
}

func minOperations(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}

	d := nums[1] - nums[0]
	if abs(d) != 1 {
		d = nums[2] - nums[1]
		if abs(d) != 1 { // 不能有多个跳变位置
			return -1
		}
	}

	p := 0
	for i := 1; i < n; i++ {
		if nums[i]-nums[i-1] != d {
			// i 是跳变位置
			if p > 0 { // 不能有多个跳变位置
				return -1
			}
			p = i
		}
	}

	if d == 1 {
		return min(p, n-p+2) // 递增
	}
	return min(p, n-p) + 1 // 递减
}

func abs(x int) int { if x < 0 { return -x }; return x }

//func minOperations(nums []int) int {
//	n := len(nums)
//	mid := slices.Index(nums, 0)
//
//	// [0, mid-1] 递增 + [mid, n-1] 递增
//	// 如果有两段，第一段的最小值必须 > 第二段的最大值
//	if slices.IsSorted(nums[:mid]) && slices.IsSorted(nums[mid:]) {
//		if mid > 0 && nums[0] < nums[n-1] {
//			return -1
//		}
//		return min(mid, n-mid+2)
//	}
//
//	// [0, mid] 递减 + [mid+1, n-1] 递减
//	// 如果有两段，第一段的最大值必须 < 第二段的最小值
//	greater := func(a, b int) int { return b - a }
//	if slices.IsSortedFunc(nums[:mid+1], greater) && slices.IsSortedFunc(nums[mid+1:], greater) {
//		if mid < n-1 && nums[0] > nums[n-1] {
//			return -1
//		}
//		return min(mid+2, n-mid+1)
//	}
//
//	return -1
//}
