package main

import (
	"math/bits"
	"slices"
)

// https://space.bilibili.com/206214
const w = bits.UintSize

type bitset []uint

// b <<= k
func (b bitset) lsh(k int) bitset {
	shift, offset := k/w, k%w
	if offset == 0 {
		// Fast path
		copy(b[shift:], b)
	} else {
		for i := len(b) - 1; i > shift; i-- {
			b[i] = b[i-shift]<<offset | b[i-shift-1]>>(w-offset)
		}
		b[shift] = b[0] << offset
	}
	clear(b[:shift])
	return b
}

// 把 >= start 的清零
func (b bitset) resetRange(start int) bitset {
	i := start / w
	b[i] &= ^(^uint(0) << (start % w))
	clear(b[i+1:])
	return b
}

// b |= c
func (b bitset) unionFrom(c bitset) {
	for i, v := range c {
		b[i] |= v
	}
}

func (b bitset) lastIndex1() int {
	for i := len(b) - 1; i >= 0; i-- {
		if b[i] != 0 {
			return i*w | (bits.Len(b[i]) - 1)
		}
	}
	return -1
}

// 167. 两数之和 II - 输入有序数组
// https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/
func twoSum(numbers []int, target int) bool {
	left, right := 0, len(numbers)-1
	for left < right {
		s := numbers[left] + numbers[right]
		if s == target {
			return true
		}
		if s > target {
			right--
		} else {
			left++
		}
	}
	return false
}

func maxTotalReward(rewardValues []int) int {
	slices.Sort(rewardValues)
	rewardValues = slices.Compact(rewardValues) // 去重

	n := len(rewardValues)
	m := rewardValues[n-1]
	if n == 1 {
		return m
	}
	if rewardValues[n-2] == m-1 || twoSum(rewardValues, m-1) {
		return m*2 - 1
	}

	f := make(bitset, m*2/w+1)
	f[0] = 1
	for _, v := range rewardValues {
		f.unionFrom(slices.Clone(f).lsh(v).resetRange(v * 2))
	}
	return f.lastIndex1()
}
