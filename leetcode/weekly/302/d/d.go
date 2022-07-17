package main

import "math"

// https://space.bilibili.com/206214/dynamic
func minOperations(nums, numsDivide []int) (ans int) {
	g := 0
	for _, x := range numsDivide {
		g = gcd(g, x)
	}
	min := math.MaxInt32
	for _, x := range nums {
		if g%x == 0 && x < min {
			min = x
		}
	}
	if min == math.MaxInt32 {
		return -1
	}
	for _, x := range nums {
		if x < min {
			ans++
		}
	}
	return
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
