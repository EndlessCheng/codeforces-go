package main

import "math/bits"

// https://space.bilibili.com/206214
func makeTheIntegerZero(num1, num2 int) int {
	for k := 1; k <= num1-num2*k; k++ {
		if k >= bits.OnesCount(uint(num1-num2*k)) {
			return k
		}
	}
	return -1
}
