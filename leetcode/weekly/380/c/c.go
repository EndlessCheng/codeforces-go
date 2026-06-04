package main

import (
	"math/bits"
	"sort"
)

// https://space.bilibili.com/206214
func findMaximumNumber(K int64, x int) int64 {
	k := int(K)
	num, pre1 := 0, 0
	for i := bits.Len(uint((k+1)<<x)) - 1; i >= 0; i-- {
		cnt := pre1<<i + i/x<<i>>1
		if cnt > k {
			continue
		}
		k -= cnt
		num |= 1 << i
		if (i+1)%x == 0 {
			pre1++
		}
	}
	return int64(num - 1)
}

func findMaximumNumber2(k int64, x int) int64 {
	ans := sort.Search(int(k+1)<<x, func(num int) bool {
		num++
		res := 0
		// 统计 [1,num] 中的第 x,2x,3x,... 个比特位上的 1 的个数
		// 注意比特位从 0 开始，不是从 1 开始，所以要减一
		for i := x - 1; num>>i > 0; i += x {
			maxPrefix := num >> (i + 1)
			// 1. prefix < maxPrefix 时，低位不受约束
			// i 位填 1，suffix 随便填
			res += maxPrefix << i
			if num>>i&1 > 0 {
				// 2. prefix = maxPrefix 且 i 位可以填 1
				// i 位填 1，suffix 可以填 [0, maxSuffix] 中的任意整数
				maxSuffix := num & (1<<i - 1)
				res += maxSuffix + 1
			}
		}
		return res > int(k)
	})
	return int64(ans)
}
