package main

import (
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func maxTotalValue(value, decay []int, m int) (ans int) {
	low := sort.Search(slices.Max(value), func(low int) bool {
		low++
		leftM := m
		for i, v := range value {
			if v >= low {
				leftM -= (v-low)/decay[i] + 1
				if leftM < 0 { // 提前跳出循环
					return false
				}
			}
		}
		return true
	})

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
