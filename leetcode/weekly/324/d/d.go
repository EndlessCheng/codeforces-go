package main

import "math/bits"

// https://space.bilibili.com/206214
func cycleLengthQueries(_ int, queries [][]int) []int {
	ans := make([]int, len(queries))
	for i, q := range queries {
		a, b := uint(q[0]), uint(q[1])
		if a > b {
			a, b = b, a
		}
		d := bits.Len(b) - bits.Len(a)
		ans[i] = d + bits.Len(b>>d^a)*2 + 1
	}
	return ans
}
