package main

import (
	"math/bits"
	"slices"
)

// https://space.bilibili.com/206214
func minSetBits1(nums []int, k int) int {
	w := bits.Len(uint(slices.Max(nums)))
	cnt := make([]int, w)
	f := make([]bool, 1<<w)
	for r, v := range nums {
		// 1. 入
		for ; v > 0; v &= v - 1 {
			cnt[bits.TrailingZeros(uint(v))]++
		}

		l := r + 1 - k
		if l < 0 { // 尚未形成第一个窗口
			continue
		}

		// 2. 计算 x
		x := 0
		for i, c := range cnt {
			if c == k {
				x |= 1 << i
			}
		}
		f[x] = true

		// 3. 出
		for v = nums[l]; v > 0; v &= v - 1 {
			cnt[bits.TrailingZeros(uint(v))]--
		}
	}

	for i := range w {
		for s := 0; s < 1<<w; s++ {
			s |= 1 << i // 优化：快速跳到 i 位是 1 的 s
			f[s] = f[s] || f[s^1<<i]
		}
	}

	mx := -1
	// 先计算 y 的补集的二进制 1 的个数的最大值
	for z, hasSubset := range f[:1<<w] {
		if !hasSubset {
			mx = max(mx, bits.OnesCount(uint(z)))
		}
	}

	if mx < 0 {
		return -1
	}
	// 最后用 w 减去 mx，即为 y 的二进制 1 的个数的最小值
	return w - mx
}

func minSetBits(nums []int, k int) int {
	w := bits.Len(uint(slices.Max(nums)))
	f := make([]bool, 1<<w)
	bottom := 0
	rightAnd := -1
	for r, v := range nums {
		rightAnd &= v
		l := r + 1 - k
		if l < 0 {
			continue
		}
		if bottom < l {
			// 重新构建一个栈
			for i := r - 1; i >= l; i-- {
				nums[i] &= nums[i+1]
			}
			bottom = r
			rightAnd = -1
		}
		f[nums[l]&rightAnd] = true
	}

	for i := range w {
		for s := 0; s < 1<<w; s++ {
			s |= 1 << i // 优化：快速跳到 i 位是 1 的 s
			f[s] = f[s] || f[s^1<<i]
		}
	}

	mx := -1
	// 先计算 y 的补集的二进制 1 的个数的最大值
	for z, hasSubset := range f[:1<<w] {
		if !hasSubset {
			mx = max(mx, bits.OnesCount(uint(z)))
		}
	}

	if mx < 0 {
		return -1
	}
	// 最后用 w 减去 mx，即为 y 的二进制 1 的个数的最小值
	return w - mx
}
