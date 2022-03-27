package main

import (
	"math"
)

// github.com/EndlessCheng/codeforces-go
func kthPalindrome(queries []int, intLength int) []int64 {
	ans := make([]int64, len(queries))
	base := int(math.Pow10((intLength - 1) / 2))
	for i, q := range queries {
		if q > 9*base {
			ans[i] = -1
			continue
		}
		v := base + q - 1 // 回文数左半部分
		x := v
		if intLength%2 == 1 { x /= 10 } // 去掉回文中心
		for ; x > 0; x /= 10 {
			v = v*10 + x%10 // 翻转 x 到 v 后
		}
		ans[i] = int64(v)
	}
	return ans
}
