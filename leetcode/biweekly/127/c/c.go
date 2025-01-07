package main

import "math"

// https://space.bilibili.com/206214
func minimumSubarrayLength(nums []int, k int) int {
	ans := math.MaxInt
	var left, bottom, rightOr int
	for right, x := range nums {
		rightOr |= x
		for left <= right && nums[left]|rightOr >= k {
			ans = min(ans, right-left+1)
			left++
			if bottom < left {
				// 重新构建一个栈
				for i := right - 1; i >= left; i-- {
					nums[i] |= nums[i+1]
				}
				bottom = right
				rightOr = 0
			}
		}
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}

func minimumSubarrayLength2(nums []int, k int) int {
	ans := math.MaxInt
	for i, x := range nums {
		if x >= k {
			return 1
		}
		// 如果 x 是 nums[j] 的子集，就退出循环
		for j := i - 1; j >= 0 && nums[j]|x != nums[j]; j-- {
			nums[j] |= x
			if nums[j] >= k {
				ans = min(ans, i-j+1)
			}
		}
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}
