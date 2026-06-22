package main

import (
	"slices"
)

// https://space.bilibili.com/206214
func maxTotalValue(value, decay []int, m int) (ans int) {
	check := func(low int) bool {
		leftM := m
		for i, v := range value {
			if v >= low {
				leftM -= (v-low)/decay[i] + 1
				if leftM < 0 { // 提前跳出循环
					return true
				}
			}
		}
		return false
	}

	low := 0
	if check(0) {
		left, right := 0, slices.Max(value)+1
		for left+1 < right {
			mid := left + (right-left)/2
			if check(mid) {
				left = mid
			} else {
				right = mid
			}
		}
		low = left
	}

	// 计算价值严格大于 low 的价值和，以及这些价值的个数
	for i, v := range value {
		if v > low {
			dec := decay[i]
			k := (v-low-1)/dec + 1
			m -= k
			ans += (v*2 - dec*(k-1)) * k
		}
	}
	ans /= 2 // 把除以 2 提到循环外面
	ans += m * low // 剩余 m 次选的价值都是 low
	return ans % 1_000_000_007
}
