package main

import "math/bits"

// github.com/EndlessCheng/codeforces-go
func Maximumweight(a, b []int, V int) int {
	mx := -1
	for i := uint(0); i < 1<<len(a); i++ {
		v, w := 0, 0
		for j := i; j > 0; j &= j - 1 {
			p := bits.TrailingZeros(j)
			v += a[p]
			w += b[p]
		}
		if v == V && w > mx {
			mx = w
		}
	}
	return mx
}
