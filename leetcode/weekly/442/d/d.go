package main

import "math/bits"

// https://space.bilibili.com/206214
// 返回 [1,n] 的单个元素的操作次数之和
func f(n int) int {
	m := bits.Len(uint(n))
	k := (m - 1) / 2 * 2
	return k<<k>>1 - 1<<k/3 + (m+1)/2*(n+1-1<<k)
}

func minOperations(queries [][]int) int64 {
	ans := 0
	for _, q := range queries {
		ans += (f(q[1]) - f(q[0]-1) + 1) / 2
	}
	return int64(ans)
}
