package main

import (
	"math/bits"
	"slices"
)

// https://space.bilibili.com/206214
func maximumAND(nums []int, k, m int) (ans int) {
	ops := make([]int, len(nums)) // 每个数的操作次数
	maxWidth := bits.Len(uint(slices.Max(nums) + k))
	for bit := maxWidth - 1; bit >= 0; bit-- {
		target := ans | 1<<bit // 注意 target 要带着 ans 已经填好的 1
		for i, x := range nums {
			j := bits.Len(uint(target &^ x))
			// j-1 是从高到低第一个 target 是 1，x 是 0 的比特位
			// target = 10110
			//      x = 11010
			//            ^
			//           j-1
			// x 高于 j-1 的比特位不变，其余变成和 target 一样
			// 上面的例子要把 010 变成 110
			mask := 1<<j - 1
			ops[i] = target&mask - x&mask
		}

		// 贪心，取前 m 小的操作次数
		slices.Sort(ops)
		sum := 0
		for _, x := range ops[:m] {
			sum += x
		}
		if sum <= k {
			ans = target // 答案的 bit 位可以填 1
		}
	}
	return
}
