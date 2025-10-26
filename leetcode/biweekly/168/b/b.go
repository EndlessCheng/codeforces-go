package main

import "strings"

// https://space.bilibili.com/206214
func maxSumOfSquares(n, sum int) string {
	if n*9 < sum {
		return ""
	}
	ans := strings.Repeat("9", sum/9)
	if sum%9 > 0 {
		ans += string('0' + byte(sum%9))
	}
	return ans + strings.Repeat("0", n-len(ans))
}
