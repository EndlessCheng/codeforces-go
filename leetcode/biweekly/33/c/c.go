package main

import "math/bits"

// github.com/EndlessCheng/codeforces-go
func minOperations(a []int) (ans int) {
	maxMul := 0
	for _, v := range a {
		ans += bits.OnesCount(uint(v))
		if bits.Len(uint(v)) > maxMul {
			maxMul = bits.Len(uint(v))
		}
	}
	return ans + maxMul - 1
}
