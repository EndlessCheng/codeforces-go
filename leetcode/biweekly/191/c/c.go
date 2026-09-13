package main

import "math"

// https://space.bilibili.com/206214
const mx = 100_001
var f [mx]int

func init() {
	f[0] = -1
	for i := 1; i < mx; i++ {
		f[i] = math.MaxInt
	}

	sum, i := 1, 1
	for sum < mx {
		for j := sum; j < mx; j++ {
			f[j] = min(f[j], f[j-sum]+i+1)
		}
		i++
		sum += i
	}
}

func minDays(n int) int {
	return f[n]
}
