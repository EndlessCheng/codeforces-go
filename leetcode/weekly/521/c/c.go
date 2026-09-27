package main

import "slices"

// https://space.bilibili.com/206214
func maxSubarray1(nums []int) (ans int) {
	mx := slices.Max(nums)
	cntS := make([]int, mx*2+1)
	cntD := make([]int, mx+1)
	left := 0

	// 枚举有效子数组的右端点为 i，那么左端点 left 最小是多少？
	for i, x := range nums {
		// x 进入窗口前，先判断：
		// 如果窗口中有两数之和等于 x，或者两数之差等于 x，那么必须缩小窗口
		for cntS[x] > 0 || cntD[x] > 0 {
			y := nums[left]
			left++
			for _, z := range nums[left:i] {
				cntS[y+z]--
				cntD[abs(y-z)]--
			}
		}

		// 元素 x 进入窗口
		for _, y := range nums[left:i] {
			cntS[x+y]++
			cntD[abs(x-y)]++
		}

		// 用子数组 [left, i] 的长度更新答案的最大值
		ans = max(ans, i-left+1)

	}

	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func maxSubarray(nums []int) (ans int) {
	mx := slices.Max(nums)
	cnt := make([]int, mx+1)
	left := 0

	// 枚举有效子数组的右端点为 i，那么左端点 left 最小是多少？
	for i, x := range nums {
		// x 进入窗口前，先判断：
		// 如果窗口中有两数之和等于 x，或者两数之差等于 x，那么必须缩小窗口
		for {
			
			
			cnt[nums[left]]--
			left++
		}

		// 元素 x 进入窗口
		cnt[x]++

		// 用子数组 [left, i] 的长度更新答案的最大值
		ans = max(ans, i-left+1)
	}

	return
}
