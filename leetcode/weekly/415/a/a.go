package main

import "math/bits"

// https://space.bilibili.com/206214
func getSneakyNumbers(nums []int) []int {
	n := len(nums) - 2
	xorAll := n ^ (n + 1) // n 和 n+1 多异或了
	for i, x := range nums {
		xorAll ^= i ^ x
	}
	shift := bits.TrailingZeros(uint(xorAll))

	ans := make([]int, 2)
	for i, x := range nums {
		if i < n {
			ans[i>>shift&1] ^= i
		}
		ans[x>>shift&1] ^= x
	}
	return ans
}
