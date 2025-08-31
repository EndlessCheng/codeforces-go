package main

import (
	"math"
	"strings"
)

// https://space.bilibili.com/206214
func minOperations(s string, k int) int {
	n := len(s)
	z := strings.Count(s, "0")
	if z == 0 {
		return 0
	}
	if n == k {
		if z == k {
			return 1
		}
		return -1
	}

	ans := math.MaxInt
	// 情况一：操作次数 m 是偶数
	if z%2 == 0 { // z 必须是偶数
		m := max((z+k-1)/k, (z+n-k-1)/(n-k)) // 下界
		ans = m + m%2                        // 把 m 往上调整为偶数
	}

	// 情况二：操作次数 m 是奇数
	if z%2 == k%2 { // z 和 k 的奇偶性必须相同
		m := max((z+k-1)/k, (n-z+n-k-1)/(n-k)) // 下界
		ans = min(ans, m|1)                    // 把 m 往上调整为奇数
	}

	if ans < math.MaxInt {
		return ans
	}
	return -1
}
