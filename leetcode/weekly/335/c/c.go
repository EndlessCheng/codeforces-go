package main

import "math/bits"

// https://space.bilibili.com/206214
func findValidSplit(nums []int) int {
	left := map[int]int{}           // left[p] 表示质数 p 首次出现的下标
	right := make([]int, len(nums)) // right[i] 表示左端点为 i 的区间的右端点的最大值
	f := func(p, i int) {
		if l, ok := left[p]; ok {
			right[l] = i // 记录左端点 l 对应的右端点的最大值
		} else {
			left[p] = i // 第一次遇到质数 p
		}
	}

	for i, x := range nums {
		if x%2 == 0 { // 单独处理 2，这样后面只需枚举奇数
			f(2, i)
			x >>= bits.TrailingZeros(uint(x))
		}
		for d := 3; d*d <= x; d += 2 { // 分解质因子
			if x%d == 0 {
				f(d, i)
				for x /= d; x%d == 0; x /= d {
				}
			}
		}
		if x > 1 {
			f(x, i)
		}
	}

	maxR := 0
	for l, r := range right {
		if l > maxR { // 最远可以到 maxR
			return maxR
		}
		maxR = max(maxR, r)
	}
	return -1
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
