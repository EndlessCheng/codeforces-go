package main

import (
	"math/bits"
	"slices"
)

// https://space.bilibili.com/206214
func maxPalindromesAfterOperations(words []string) (ans int) {
	tot, mask := 0, 0
	for _, w := range words {
		tot += len(w)
		for _, c := range w {
			mask ^= 1 << (c - 'a')
		}
	}
	tot -= bits.OnesCount(uint(mask)) // 减去出现次数为奇数的字母

	slices.SortFunc(words, func(a, b string) int { return len(a) - len(b) })
	for _, w := range words {
		tot -= len(w) / 2 * 2 // 偶数不变，奇数减一
		if tot < 0 {
			break
		}
		ans++
	}
	return
}
