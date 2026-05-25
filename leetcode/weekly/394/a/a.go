package main

import "math/bits"

// https://space.bilibili.com/206214
func numberOfSpecialChars(word string) int {
	mask := [2]int{} // 大写字母集合、小写字母集合
	for _, c := range word {
		// 用 c>>5&1 区分大小写，c&31 获取 c 是第几个字母
		mask[c>>5&1] |= 1 << (c & 31)
	}
	return bits.OnesCount(uint(mask[0] & mask[1])) // 计算交集大小
}
