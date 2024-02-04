package main

import "slices"

// https://space.bilibili.com/206214
func triangleType(nums []int) string {
	slices.Sort(nums)
	x, y, z := nums[0], nums[1], nums[2]
	if x+y <= z { // 排序后，只需比较 x+y 和 z
		return "none"
	}
	if x == z { // 排序了，说明 y 也和 x z 相等
		return "equilateral"
	}
	if x == y || y == z {
		return "isosceles"
	}
	return "scalene"
}
