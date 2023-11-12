package main

import (
	"math/bits"
	"slices"
)

// https://space.bilibili.com/206214
func maximumStrongPairXor(nums []int) (ans int) {
	slices.Sort(nums)
	highBit := bits.Len(uint(nums[len(nums)-1])) - 1
	mp := map[int]int{}
	mask := 0
	for i := highBit; i >= 0; i-- { // 从最高位开始枚举
		clear(mp)
		mask |= 1 << i
		newAns := ans | 1<<i // 这个比特位可以是 1 吗？
		for _, y := range nums {
			maskY := y & mask // 低于 i 的比特位置为 0
			if x, ok := mp[newAns^maskY]; ok && x*2 >= y {
				ans = newAns // 这个比特位可以是 1
				break
			}
			mp[maskY] = y
		}
	}
	return
}
