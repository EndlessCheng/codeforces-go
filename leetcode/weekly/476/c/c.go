package main

import (
	"math"
	"strconv"
)

// https://space.bilibili.com/206214
func countDistinct(n int64) int64 {
	s := strconv.FormatInt(n, 10)
	m := len(s)

	// 计算长度小于 m 的不含 0 的整数个数
	// 9 + 9^9 + ... + 9^(m-1) = (9^m - 9) / 8
	pow9 := int64(math.Pow(9, float64(m)))
	ans := (pow9 - 9) / 8

	// 计算长度恰好等于 m 的不含 0 的整数个数
	for i, d := range s {
		if d == '0' { // 只能填 0，不合法，跳出循环
			break
		}
		// 这一位填 1 到 d-1，后面的数位可以随便填 1 到 9
		v := d - '1'
		if i == m-1 {
			v++ // 最后一位可以等于 d
		}
		pow9 /= 9
		ans += int64(v) * pow9
		// 然后，这一位填 d，继续遍历
	}
	return ans
}
