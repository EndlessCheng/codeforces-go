package main

import "math"

// https://space.bilibili.com/206214
func maxSubarraySum1(nums []int, k int) int64 {
	solve := func(isMul bool) int64 {
		res := int64(math.MinInt)
		n := len(nums)
		// f[i+1][0] 表示右端点为 i 的最大子数组和，且不修改任何元素
		// f[i+1][1] 表示右端点为 i 的最大子数组和，且修改了 nums[i]
		// f[i+1][2] 表示右端点为 i 的最大子数组和，且在 nums[i] 的左边发生了修改（没有修改 nums[i]）
		f := make([][3]int64, n+1)
		for i, x := range nums {
			x := int64(x)
			y := x
			if isMul {
				y *= int64(k)
			} else {
				y /= int64(k)
			}
			// 不修改 x，和 f[i][0] 拼起来，或者 x 是子数组的第一个数
			f[i+1][0] = max(f[i][0], 0) + x
			// 修改 x，和 f[i][0] 或者 f[i][1] 拼起来，或者 y 是子数组的第一个数
			f[i+1][1] = max(f[i][0], f[i][1], 0) + y
			// 不修改 x，和 f[i][1] 或者 f[i][2] 拼起来
			f[i+1][2] = max(f[i][1], f[i][2]) + x
			// 枚举子数组的右端点为 i
			res = max(res, f[i+1][1], f[i+1][2])
		}
		return res
	}
	return max(solve(true), solve(false))
}

func maxSubarraySum(nums []int, k int) int64 {
	solve := func(isMul bool) int64 {
		res := int64(math.MinInt)
		var f0, f1, f2 int64
		for _, x := range nums {
			x := int64(x)
			y := x
			if isMul {
				y *= int64(k)
			} else {
				y /= int64(k)
			}
			f2 = max(f1, f2) + x
			f1 = max(f0, f1, 0) + y
			f0 = max(f0, 0) + x
			res = max(res, f1, f2)
		}
		return res
	}
	return max(solve(true), solve(false))
}
