package main

import (
	"math/bits"
	"slices"
)

// https://space.bilibili.com/206214
func maxProduct1(nums []int) int64 {
	w := bits.Len(uint(slices.Max(nums)))
	u := 1 << w
	f := make([]int, u)
	for _, x := range nums {
		f[x] = x
	}

	for s := 3; s < u; s++ { // 从小到大枚举集合 s
		for i := range w { // 枚举 s 中的元素 i
			if s>>i&1 > 0 { // i 属于集合 s
				f[s] = max(f[s], f[s^1<<i]) // 从 s 的子集 s \ {i} 转移过来
			}
		}
	}

	ans := 0
	for _, x := range nums {
		ans = max(ans, x*f[u-1^x])
	}
	return int64(ans)
}

func maxProduct(nums []int) int64 {
	w := bits.Len(uint(slices.Max(nums)))
	u := 1 << w
	f := make([]int, u)
	for _, x := range nums {
		f[x] = x
	}

	for i := range w {
		for s := 3; s < u; s++ {
			s |= 1 << i // 快速跳到第 i 位是 1 的 j
			f[s] = max(f[s], f[s^1<<i])
		}
	}

	ans := 0
	for _, x := range nums {
		ans = max(ans, x*f[u-1^x])
	}
	return int64(ans)
}
