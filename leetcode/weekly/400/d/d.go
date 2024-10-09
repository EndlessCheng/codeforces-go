package main

import "math"

// https://space.bilibili.com/206214
func minimumDifference(nums []int, k int) int {
	ans := math.MaxInt
	var left, bottom, rightOr int
	for right, x := range nums {
		rightOr |= x
		for left <= right && nums[left]|rightOr > k {
			ans = min(ans, (nums[left]|rightOr)-k)
			if bottom <= left {
				// 重新构建一个栈
				// 由于 left 即将移出窗口，只需计算到 left+1
				for i := right - 1; i > left; i-- {
					nums[i] |= nums[i+1]
				}
				bottom = right
				rightOr = 0
			}
			left++
		}
		if left <= right {
			ans = min(ans, k-(nums[left]|rightOr))
		}
	}
	return ans
}

func minimumDifference2(nums []int, k int) int {
	ans := math.MaxInt
	for i, x := range nums {
		ans = min(ans, abs(x-k))
		for j := i - 1; j >= 0 && nums[j]|x != nums[j]; j-- {
			nums[j] |= x
			ans = min(ans, abs(nums[j]-k))
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
