package main

import "math"

// https://space.bilibili.com/206214
func maxScore(a, b []int) int64 {
	f := [5]int64{}
	for j := 1; j < 5; j++ {
		f[j] = math.MinInt64 / 2
	}
	for _, y := range b {
		for j := 3; j >= 0; j-- {
			f[j+1] = max(f[j+1], f[j]+int64(a[j])*int64(y))
		}
	}
	return f[4]
}
