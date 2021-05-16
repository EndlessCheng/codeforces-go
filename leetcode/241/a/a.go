package main

import "math/bits"

// github.com/EndlessCheng/codeforces-go
func subsetXORSum(a []int) (ans int) {
	for sub := uint(0); sub < 1<<len(a); sub++ {
		xor := 0
		for s := sub; s > 0; s &= s - 1 {
			xor ^= a[bits.TrailingZeros(s)]
		}
		ans += xor
	}
	return
}

// github.com/EndlessCheng/codeforces-go
func subsetXORSum2(a []int) (ans int) {
	sum := make([]int, 1<<len(a))
	for p, v := range a {
		for s := 0; s < 1<<p; s++ {
			sum[1<<p|s] = sum[s] ^ v
			ans += sum[1<<p|s]
		}
	}
	return
}
