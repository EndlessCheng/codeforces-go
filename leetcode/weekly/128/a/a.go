package main

import "math/bits"

// github.com/EndlessCheng/codeforces-go
func bitwiseComplement(n int) int {
	if n == 0 {
		return 1
	}
	w := bits.Len(uint(n))
	return 1<<w - 1 ^ n
}
