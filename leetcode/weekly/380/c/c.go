package main

import "math/bits"

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
