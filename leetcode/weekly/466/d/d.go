package main

import (
	"math/bits"
	"strconv"
)

// https://space.bilibili.com/206214
func countBinaryPalindromes1(n int64) int {
	s := strconv.FormatUint(uint64(n), 2)
	m := len(s)

	// 二进制长度小于 m，随便填
	ans := 1
	// 枚举二进制长度，最高位填 1，回文数左半的其余位置随便填
	for i := 1; i < m; i++ {
		ans += 1 << ((i - 1) / 2)
	}

	var dfs func(int, int, bool) int
	dfs = func(i, pal int, limit bool) (res int) {
		if !limit {
			// 回文数左半的其余位置随便填
			return 1 << ((m+1)/2 - i)
		}

		if i == (m+1)/2 {
			// 左半反转到右半
			// 如果 m 是奇数，那么去掉回文中心再反转
			for v := pal >> (m % 2); v > 0; v /= 2 {
				pal = pal*2 + v%2
			}
			if pal > int(n) {
				return 0
			}
			return 1
		}

		up := int(s[i] - '0')
		for d := 0; d <= up; d++ {
			res += dfs(i+1, pal*2+d, limit && d == up)
		}
		return res
	}

	// 最高位一定是 1，从 i=1 开始填
	return ans + dfs(1, 1, true)
}

func countBinaryPalindromes(n int64) int {
	if n == 0 {
		return 1
	}

	m := bits.Len(uint(n))
	k := (m - 1) / 2

	// 二进制长度小于 m
	ans := 2<<k - 1
	if m%2 == 0 {
		ans += 1 << k
	}

	// 二进制长度等于 m，且回文数的左半小于 n 的左半
	left := n >> (m / 2)
	ans += int(left) - 1<<k

	// 二进制长度等于 m，且回文数的左半等于 n 的左半
	right := bits.Reverse32(uint32(left>>(m%2))) >> (32 - m/2)
	if left<<(m/2)|int64(right) <= n {
		ans++
	}

	return ans
}
