package main

import "slices"

// https://space.bilibili.com/206214
func smallestSubarrays1(nums []int) []int {
	ans := make([]int, len(nums))
	for i, x := range nums { // 计算右端点为 i 的子数组的或值
		ans[i] = 1 // 子数组的长度至少是 1
		// 循环直到 nums[j] 无法增大，其左侧元素也无法增大
		for j := i - 1; j >= 0 && nums[j]|x != nums[j]; j-- {
			nums[j] |= x       // nums[j] 增大，现在 nums[j] = 原数组 nums[j] 到 nums[i] 的或值
			ans[j] = i - j + 1 // nums[j] 最后一次增大时的子数组长度就是答案
		}
	}
	return ans
}

func smallestSubarrays(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	leftOr, sufOr := 0, 0
	right, bottom := n-1, n-1
	for left, x := range slices.Backward(nums) {
		sufOr |= x
		leftOr |= x
		for right >= left && leftOr|nums[right] == sufOr {
			right--
			// 栈为空
			if bottom > right {
				// 重新构建一个栈，栈底为 left，栈顶为 right
				for i := left + 1; i <= right; i++ {
					nums[i] |= nums[i-1]
				}
				bottom = left
				leftOr = 0
			}
		}
		// 循环结束后 [left,right] 不满足要求，但 [left, right+1] 满足要求
		ans[left] = right - left + 2
	}
	return ans
}
