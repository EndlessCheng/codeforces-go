package main

import (
	"math"
)

// https://space.bilibili.com/206214
func maximumEnergy(energy []int, k int) int {
	n := len(energy)
	ans := math.MinInt
	for i := n - k; i < n; i++ {
		s := 0
		for j := i; j >= 0; j -= k {
			s += energy[j]
			ans = max(ans, s)
		}
	}
	return ans
}
