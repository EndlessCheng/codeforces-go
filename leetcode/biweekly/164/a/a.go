package main

import "math"

// https://space.bilibili.com/206214
func getLeastFrequentDigit(n int) (ans int) {
	// 统计每个数字的出现次数
	cnt := [10]int{}
	for n > 0 {
		cnt[n%10]++
		n /= 10
	}

	// 找出现次数最小的数字
	minC := math.MaxInt
	for i, c := range cnt {
		if c > 0 && c < minC {
			minC = c
			ans = i
		}
	}
	return
}
