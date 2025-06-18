package main

import (
	"math"
	"slices"
)

// https://space.bilibili.com/206214/dynamic
func partitionArray(nums []int, k int) (ans int) {
	slices.Sort(nums)
	mn := math.MinInt / 2 // 防止减法溢出
	for _, x := range nums {
		if x-mn > k { // 必须分割
			ans++
			mn = x // mn 是下一段的最小值
		}
	}
	return
}
