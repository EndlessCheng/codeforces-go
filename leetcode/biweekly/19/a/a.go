package main

import "math/bits"

// github.com/EndlessCheng/codeforces-go
func numberOfSteps(n int) (ans int) {
	if n == 0 {
		return
	}
	return bits.Len(uint(n)) - 1 + bits.OnesCount(uint(n))
}
