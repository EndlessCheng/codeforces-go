package main

import "math"

// https://space.bilibili.com/206214/
func maximumSum(nums []int) int {
	ans := -1
	mx := [82]int{} // 至多 9 个 9 相加
	for i := range mx {
		mx[i] = math.MinInt
	}
	for _, num := range nums { // 枚举 num = nums[j]
		s := 0 // num 的数位和
		for x := num; x > 0; x /= 10 { // 枚举 num 的每个数位
			s += x % 10
		}
		ans = max(ans, mx[s]+num) // 左边找一个数位和也为 s 的最大的 nums[i]
		mx[s] = max(mx[s], num)   // 维护数位和等于 s 的最大元素
	}
	return ans
}
