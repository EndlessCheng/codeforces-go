package main

import (
	"math"
)

// https://space.bilibili.com/206214
func countCommas(n int64) (ans int64) {
	// 从右往左，枚举逗号的位置
	for low := int64(1000); low <= n; low *= 1000 {
		// [low, n] 中的每个数都在这个位置上有一个逗号
		ans += n - low + 1
	}
	return
}

func countCommas2(n int64) (ans int64) {
	k := 5 // n == 1e15 时 Log10 有误差，需要特判
	if n < 1e15 {
		k = int(math.Log10(float64(n))) / 3
	}
	return int64(k)*(n+1) - (int64(math.Pow10(k*3+3))-1000)/999
}
