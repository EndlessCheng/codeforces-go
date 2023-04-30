package main

import "math/bits"

// https://space.bilibili.com/206214
func findThePrefixCommonArray(a, b []int) []int {
	ans := make([]int, len(a))
	var p, q uint
	for i, x := range a {
		p |= 1 << x
		q |= 1 << b[i]
		ans[i] = bits.OnesCount(p & q)
	}
	return ans
}
