package main

import (
	"math/bits"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func findKthSmallest(coins []int, k int) int64 {
	ans := sort.Search(slices.Max(coins)*k, func(m int) bool {
		cnt := 0
	next:
		for i := uint(1); i < 1<<len(coins); i++ { // 枚举所有非空子集
			lcmRes := 1 // 计算子集 LCM
			for j := i; j > 0; j &= j - 1 {
				lcmRes = lcm(lcmRes, coins[bits.TrailingZeros(j)])
				if lcmRes > m { // 太大了
					continue next
				}
			}
			c := m / lcmRes
			if bits.OnesCount(i)%2 == 0 {
				c = -c
			}
			cnt += c
		}
		return cnt >= k
	})
	return int64(ans)
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}
