package main

import (
	"math"
	"slices"
)

// https://space.bilibili.com/206214
func maximumEnergy1(energy []int, k int) int {
	n := len(energy)
	ans := math.MinInt
	for i := n - k; i < n; i++ { // 枚举终点 i
		sufSum := 0
		for j := i; j >= 0; j -= k {
			sufSum += energy[j] // 计算后缀和
			ans = max(ans, sufSum)
		}
	}
	return ans
}

func maximumEnergy(energy []int, k int) int {
	for i := len(energy) - k - 1; i >= 0; i-- {
		energy[i] += energy[i+k]
	}
	return slices.Max(energy)
}
