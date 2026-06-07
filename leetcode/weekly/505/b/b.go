package main

import (
	"math"
	"math/bits"
)

// https://space.bilibili.com/206214
var cost [1 << 12]int

func init() {
	for x := 1; x < len(cost); x++ {
		if x&(x>>1) > 0 { // 有两个连续的 1
			cost[x] = math.MaxInt // 不满足要求
		} else {
			// 去掉 x 中的一个比特位（最低位还是最高位都可以），计算 DP
			cost[x] = cost[x&(x-1)] + bits.TrailingZeros(uint(x))
		}
	}
}

func generateValidStrings(n, k int) (ans []string) {
	s := make([]byte, n)
	for x, c := range cost[:1<<n] {
		if c > k {
			continue
		}
		for j := range s {
			s[j] = '0' + byte(x&1)
			x >>= 1
		}
		ans = append(ans, string(s))
	}
	return
}
